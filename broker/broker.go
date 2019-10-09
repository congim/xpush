package broker

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/congim/xpush/broker/internal/cluster"
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/network/listener"
	"github.com/congim/xpush/pkg/network/websocket"
	"github.com/kelindar/tcp"
	"go.uber.org/zap"
)

type Broker struct {
	cluster  cluster.Cluster
	conf     *config.Config
	http     *http.Server
	tcp      *tcp.Server
	logger   *zap.Logger
	listener *listener.Listener
	// protocol
	// storage
	// verify
	// channel
}

var gBroker *Broker

// New return broker struct
func New(conf *config.Config, logger *zap.Logger) *Broker {
	gBroker = &Broker{
		conf:   conf,
		http:   new(http.Server),
		tcp:    new(tcp.Server),
		logger: logger,
	}

	if gBroker.conf.Cluster != nil {
		gBroker.cluster = cluster.New(gBroker.conf.Cluster, gBroker.logger, notify)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", gBroker.onRequest)
	gBroker.http.Handler = mux
	gBroker.tcp.OnAccept = func(conn net.Conn) {
		c := newConn(conn, gBroker)
		go c.Process()
	}

	return gBroker
}

// Start server
func (b *Broker) Start() error {
	// 存储启动
	// 通道服务初始化
	// 网络监听
	// 初始化集群
	if err := b.ClusterStart(); err != nil {
		b.logger.Error("cluster start", zap.Error(err))
		return err
	}

	//errChan := make(chan error, 1)
	//go b.listen(errChan)
	//
	//go func() {
	//	select {
	//	case err, ok := <-errChan:
	//		if ok {
	//			b.logger.Fatal("start", zap.Error(err))
	//		}
	//	}
	//}()
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
		c := newConn(conn, b)
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
