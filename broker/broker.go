package broker

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/congim/xpush/broker/internal/cache"
	"github.com/congim/xpush/broker/internal/cluster"
	"github.com/congim/xpush/broker/internal/uid"
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"github.com/congim/xpush/pkg/network/listener"
	"github.com/congim/xpush/pkg/network/websocket"
	"github.com/congim/xpush/provider/storage"
	"github.com/kelindar/tcp"
	"go.uber.org/zap"
)

var logger *zap.Logger

type Broker struct {
	cluster  cluster.Cluster
	conf     *config.Config
	http     *http.Server
	tcp      *tcp.Server
	logger   *zap.Logger
	listener *listener.Listener
	topics   sync.Map
	storage  storage.Storage
	uid      uid.UIDs
	cache    cache.Cache

	// protocol
	// verify
	// channel
}

var gBroker *Broker

// New return broker struct
func New(conf *config.Config, l *zap.Logger) *Broker {

	// test start
	msg := message.New()
	msg.Type = message.MsgPub
	msg.Version = 1
	msg.Topic = "test"
	msg.ID = "123456"
	msg.Payload = []byte("hello xpush !")
	b, _ := msg.Encode()
	log.Println("body", string(b))
	// test end

	gBroker = &Broker{
		conf:    conf,
		http:    new(http.Server),
		tcp:     new(tcp.Server),
		logger:  l,
		storage: storage.New(conf.Storage, l),
		uid:     uid.New(),
		cache:   cache.New(conf.Cache, l),
	}

	logger = l

	if gBroker.conf.Cluster != nil {
		gBroker.cluster = cluster.New(gBroker.conf.Cluster, gBroker.logger, notify)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", gBroker.onRequest)
	gBroker.http.Handler = mux
	gBroker.tcp.OnAccept = func(conn net.Conn) {
		c := newConn(conn, gBroker, gBroker.conf.Listener.ReadTimeOut)
		go c.Process()
	}

	return gBroker
}

// Start server
func (b *Broker) Start() error {
	if err := b.storage.Init(); err != nil {
		b.logger.Error("storage init", zap.Error(err))
		return err
	}
	// 初始化集群
	if err := b.ClusterStart(); err != nil {
		b.logger.Error("cluster start", zap.Error(err))
		return err
	}

	errChan := make(chan error, 1)
	go b.listen(errChan)

	go func() {
		select {
		case err, ok := <-errChan:
			if ok {
				b.logger.Fatal("start", zap.Error(err))
			}
		}
	}()
	return nil
}

func (b *Broker) ClusterStart() error {
	if err := b.cluster.Start(); err != nil {
		b.logger.Error("cluster start", zap.Error(err))
		return err
	}

	if err := b.cluster.Join(); err != nil {
		b.logger.Error("cluster start", zap.Error(err))
		return err
	}
	return nil
}

// Close server
func (b *Broker) Close() {
	if b.cluster != nil {
		_ = b.cluster.Close()
	}
	if b.listener != nil {
		_ = b.listener.Close()
	}
}

func (b *Broker) onRequest(w http.ResponseWriter, r *http.Request) {
	if conn, ok := websocket.TryUpgrade(w, r); ok {
		c := newConn(conn, b, gBroker.conf.Listener.ReadTimeOut)
		go c.Process()
	}
}

// Listen starts the service.
func (b *Broker) listen(errChan chan<- error) {
	var err error
	tlsConf := &tls.Config{}
	if b.conf.Listener.IsTLS {
		cer, err := tls.LoadX509KeyPair(b.conf.Listener.Certificate, b.conf.Listener.PrivateKey)
		if err != nil {
			b.logger.Error("TLS", zap.Error(err), zap.String("tls_pem", b.conf.Listener.Certificate), zap.String("tls_key", b.conf.Listener.PrivateKey))
			errChan <- err
			return
		}
		tlsConf.Certificates = append(tlsConf.Certificates, cer)
	} else {
		tlsConf = nil
	}

	b.listener, err = listener.New(b.conf.Listener.ListenAddr, tlsConf)
	if err != nil {
		b.logger.Error("create listener err", zap.Error(err))
		errChan <- err
		return
	}

	// Set the read timeout on our mux listener
	b.listener.SetReadTimeout(120 * time.Second)

	// Configure the matchers
	b.listener.ServeAsync(listener.MatchHTTP(), b.http.Serve)
	b.listener.ServeAsync(listener.MatchAny(), b.tcp.Serve)
	err = b.listener.Serve()
	if err != nil {
		b.logger.Error("listener err", zap.Error(err))
		errChan <- err
		return
	}
	return
}

func (b *Broker) subscribe(topic string, cid uint64, con *Conn) error {
	conns, ok := b.topics.Load(topic)
	if !ok {
		conns = new(sync.Map)
		b.topics.Store(topic, conns)
	}
	// 这里cid本服务内自增，所以不需要查询是否存在再删除，直接保存即可
	conns.(*sync.Map).Store(cid, con)
	return nil
}

func (b *Broker) pushOnline(owner uint64, msg *message.Message) error {
	conns, ok := b.topics.Load(msg.Topic)
	if !ok {
		return nil
	}

	conns.(*sync.Map).Range(func(cid, conn interface{}) bool {
		if cid != owner {
			if err := conn.(*Conn).Publish(msg); err != nil {
				logger.Warn("push failed", zap.Uint64("cid", cid.(uint64)), zap.String("topic", msg.Topic), zap.Error(err))
			}
		}
		return true
	})

	return nil
}

func (b *Broker) pushOnlineWithoutOwner(msg *message.Message) error {
	conns, ok := b.topics.Load(msg.Topic)
	if !ok {
		return nil
	}

	conns.(*sync.Map).Range(func(cid, conn interface{}) bool {
		if err := conn.(*Conn).Publish(msg); err != nil {
			logger.Warn("push failed", zap.Uint64("cid", cid.(uint64)), zap.String("topic", msg.Topic), zap.Error(err))
		}
		return true
	})

	// @TODO 记录msgID to mqttID的映射关系，等待Ack的时候记录已读

	return nil
}
