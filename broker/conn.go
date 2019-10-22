package broker

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/congim/xpush/pkg/message"
	"github.com/congim/xpush/pkg/network/mqtt"
	"go.uber.org/zap"
)

type Conn struct {
	broker    *Broker
	socket    net.Conn // The transport used to read and write messages.
	username  string
	password  string
	keepalive uint16
	clientID  string
	cid       uint64
	msgID     uint16
	msgIDLock sync.Mutex
	msgQueue  chan []*message.Message
}

func (c *Conn) getMsgID() (uint16, error) {
	c.msgIDLock.Lock()
	defer c.msgIDLock.Unlock()
	if c.msgID >= 65535 {
		//return 0, fmt.Errorf("msgid exception")
		c.msgID = 0
	}
	c.msgID++
	return c.msgID, nil
}

func newConn(conn net.Conn, broker *Broker, readTout uint16) *Conn {
	return &Conn{
		socket:    conn,
		broker:    broker,
		keepalive: readTout,
		msgQueue:  make(chan []*message.Message, 500),
	}
}

// Process processes the messages.
func (c *Conn) Process() {
	defer func() {
		if err := recover(); err != nil {
			logger.Warn("process", zap.Any("recover", err))
		}
	}()
	stopC := make(chan struct{}, 1)
	defer func() {
		_ = c.Close()
		close(stopC)
	}()

	go c.sendLoop(stopC)
	c.readLoop()
}

func (c *Conn) sendLoop(stopC <-chan struct{}) {
	for {
		select {
		case <-stopC:
			return
		case msgs, ok := <-c.msgQueue:
			if ok {
				for _, msg := range msgs {
					if err := c.publish(msg); err != nil {
						logger.Warn("pushlish failed", zap.Uint64("cid", c.cid), zap.String("topic", msg.Topic), zap.Error(err))
					}
				}
			}
			break
		}
	}
}

func (c *Conn) readLoop() {
	maxSize := c.broker.conf.Limit.MessageSize
	reader := bufio.NewReaderSize(c.socket, 65536)
	for {
		// Set read/write deadlines so we can close dangling connections
		_ = c.socket.SetDeadline(time.Now().Add(time.Second * time.Duration(c.keepalive)))
		//if c.limit.Limit() {
		//	time.Sleep(50 * time.Millisecond)
		//	continue
		//}

		// Decode an incoming MQTT packet
		msg, err := mqtt.DecodePacket(reader, maxSize)
		if err != nil {
			return
		}

		// Handle the receive
		if err := c.onReceive(msg); err != nil {
			return
		}
	}
}

// Close terminates the connection.
func (c *Conn) Close() error {
	if r := recover(); r != nil {
		//logging.LogAction("closing", fmt.Sprintf("panic recovered: %s \n %s", r, debug.Stack()))
	}

	// Unsubscribe from everything, no need to lock since each Unsubscribe is
	// already locked. Locking the 'Close()' would result in a deadlock.
	//for _, counter := range c.subs.All() {
	//	c.service.onUnsubscribe(counter.Ssid, c)
	//	c.service.notifyUnsubscribe(c, counter.Ssid, counter.Channel)
	//}

	// Close the transport and decrement the connection counter
	//atomic.AddInt64(&c.service.connections, -1)
	//logging.LogTarget("conn", "closed", c.guid)
	return c.socket.Close()
}

// onReceive handles an MQTT receive.
func (c *Conn) onReceive(msg mqtt.Message) error {
	//defer c.MeasureElapsed("rcv."+msg.String(), time.Now())
	switch msg.Type() {
	// We got an attempt to connect to MQTT.
	case mqtt.TypeOfConnect:
		var result uint8
		if !c.onConnect(msg.(*mqtt.Connect)) {
			result = 0x05 // Unauthorized
		}
		// Write the ack
		ack := mqtt.Connack{ReturnCode: result}
		if _, err := ack.EncodeTo(c.socket); err != nil {
			return err
		}

	// We got an attempt to subscribe to a channel.
	case mqtt.TypeOfSubscribe:
		packet := msg.(*mqtt.Subscribe)
		ack := mqtt.Suback{
			MessageID: packet.MessageID,
			Qos:       make([]uint8, 0, len(packet.Subscriptions)),
		}
		// @TODO 订阅处理
		// Subscribe for each subscription
		for _, sub := range packet.Subscriptions {
			if err := c.onSubscribe(string(sub.Topic)); err != nil {
				ack.Qos = append(ack.Qos, 0x80) // 0x80 indicate subscription failure
				//c.notifyError(err, packet.MessageID)
				continue
			}
			// Append the QoS
			ack.Qos = append(ack.Qos, sub.Qos)
		}

		// Acknowledge the subscription
		if _, err := ack.EncodeTo(c.socket); err != nil {
			return err
		}

	// We got an attempt to unsubscribe from a channel.
	case mqtt.TypeOfUnsubscribe:
		packet := msg.(*mqtt.Unsubscribe)
		ack := mqtt.Unsuback{MessageID: packet.MessageID}
		// @TODO 取消订阅
		// Unsubscribe from each subscription
		//for _, sub := range packet.Topics {
		//	if err := c.onUnsubscribe(sub.Topic); err != nil {
		//		c.notifyError(err, packet.MessageID)
		//	}
		//}

		// Acknowledge the unsubscription
		if _, err := ack.EncodeTo(c.socket); err != nil {
			return err
		}

	// We got an MQTT ping response, respond appropriately.
	case mqtt.TypeOfPingreq:
		ack := mqtt.Pingresp{}
		if _, err := ack.EncodeTo(c.socket); err != nil {
			return err
		}

	case mqtt.TypeOfDisconnect:
		// @TODO 清理缓存等信息
		return nil

	case mqtt.TypeOfPublish:
		packet := msg.(*mqtt.Publish)
		// @TODO 优化错误处理
		msg := message.New()
		if err := msg.Decode(packet.Payload); err != nil {
			logger.Warn("msg decode failed", zap.Error(err))
			return err
		}
		if err := c.onPublish(packet, msg); err != nil {
			logger.Warn("onPublish failed", zap.Uint64("cid", c.cid), zap.String("userName", c.username), zap.Error(err))
		}

		// 计数器
		// Acknowledge the publication
		if packet.Header.QOS > 0 {
			ack := mqtt.Puback{
				MessageID: packet.MessageID,
			}
			if _, err := ack.EncodeTo(c.socket); err != nil {
				return err
			}
		}
	case mqtt.TypeOfPuback:

		break
	}
	return nil
}

// onConnect handles the connection authorization
func (c *Conn) onConnect(packet *mqtt.Connect) bool {
	// @TODO 账号密码校验
	c.username = string(packet.Username)
	c.password = string(packet.Password)
	c.clientID = string(packet.ClientID)
	if 0 < packet.KeepAlive && c.keepalive < packet.KeepAlive {
		c.keepalive = packet.KeepAlive
	}

	// 申请cid
	c.cid = c.broker.uid.Uid()

	// 缓存中存储用户登陆信息
	if err := c.broker.cache.Login(c.cid, c.broker.conf.Cluster.Name); err != nil {
		logger.Warn("cache login failed", zap.String("userName", c.username), zap.String("clientID", c.clientID), zap.Uint64("cid", c.cid), zap.Error(err))
		return false
	}
	return true
}

func (c *Conn) onSubscribe(topic string) error {
	// @TODO 数据库中存储订阅主题信息
	if err := c.broker.cache.Subscribe(c.username, topic); err != nil {
		logger.Warn("subscribe cache failed", zap.Uint64("cid", c.cid), zap.String("userName", c.username), zap.String("topic", topic), zap.Error(err))
		return err
	}

	// @TODO 建立topic和clientID直接映射关系
	if err := c.broker.subscribe(topic, c.cid, c); err != nil {
		logger.Warn("subscribe failed", zap.String("userName", c.username), zap.String("clientID", c.clientID), zap.Uint64("cid", c.cid), zap.Error(err))
		return err
	}
	return nil
}

func (c *Conn) onPublish(packet *mqtt.Publish, msg *message.Message) error {
	switch msg.Type {
	case message.MsgPub:
		if err := c.broker.storage.Store([]*message.Message{msg}); err != nil {
			logger.Warn("store msg failed", zap.Error(err))
			return err
		}

		// @TODO 检测本机在线用并推送
		if err := c.broker.pushOnline(c.cid, msg); err != nil {
			logger.Warn("push online failed", zap.Error(err))
			//return err
		}

		// @TODO 计数器更新
		if err := c.broker.cache.PubCount(msg.Topic, 1); err != nil {
			logger.Warn("publish count failed", zap.Error(err))
		}

		// @TODO 将消息推送到其他集群上
		_, _ = c.broker.cluster.OnAllMessage(msg)
		//log.Println(replys)

		break
	case message.MsgPull:

		break
	default:
		return fmt.Errorf("unknow msg type, type is %d", msg.Type)
	}

	// 消息存储
	// 集群同步&&推送消息
	// 计数器
	return nil
}

func (c *Conn) Publish(packet *message.Message) error {
	c.msgQueue <- []*message.Message{packet}
	return nil
}

func (c *Conn) publish(packet *message.Message) error {
	payload, err := packet.Encode()
	if err != nil {
		logger.Warn("packet encode failed", zap.Uint64("cid", c.cid), zap.String("topic", packet.Topic), zap.Error(err))
		return err
	}

	msgID, _ := c.getMsgID()
	msg := mqtt.Publish{
		Header: &mqtt.StaticHeader{
			QOS: 0,
		},
		Topic:     []byte(packet.Topic),
		MessageID: msgID,
		Payload:   payload,
	}

	_, err = msg.EncodeTo(c.socket)
	return err
}
