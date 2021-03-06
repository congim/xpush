package broker

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/congim/xpush/pkg/message"
	"github.com/congim/xpush/pkg/network/mqtt"
	"go.uber.org/zap"
)

// Conn conn
type Conn struct {
	broker     *Broker
	cid        uint64                  // conn id
	socket     net.Conn                // The transport used to read and write messages.
	username   string                  // username
	password   string                  // password
	clientID   string                  // clientID
	keepalive  uint16                  // keepalive
	mqttIDLock sync.Mutex              // mqttIDLock
	sendQueue  chan []*message.Message // 发送队列
	pubIDs     sync.Map                // 发送消息id缓存,ack之后删除
	topics     sync.Map                // topics
	id         uint16                  // mqtt id
}

// mqttID 自增ID，超出长度回滚
func (c *Conn) mqttID() (uint16, error) {
	c.mqttIDLock.Lock()
	if c.id >= 65535 {
		c.id = 0
	}
	c.id++
	c.mqttIDLock.Unlock()

	return c.id, nil
}

func newConn(conn net.Conn, broker *Broker, readTout uint16) *Conn {
	return &Conn{
		socket:    conn,
		broker:    broker,
		keepalive: readTout,
		sendQueue: make(chan []*message.Message, 500),
	}
}

// Process processes the messages.
func (c *Conn) Process() {
	defer func() {
		//if err := recover(); err != nil {
		//	logger.Warn("process", zap.Any("recover", err))
		//}
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
		case msgs, ok := <-c.sendQueue:
			if ok {
				if err := c.publish(msgs); err != nil {
					logger.Warn("pushlish failed", zap.Uint64("cid", c.cid), zap.Error(err))
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
		if err := c.socket.SetDeadline(time.Now().Add(time.Second * time.Duration(c.keepalive))); err != nil {
			logger.Error("setDeadline failed", zap.Error(err))
			return
		}
		//if c.limit.Limit() {
		//	time.Sleep(50 * time.Millisecond)
		//	continue
		//}

		// Decode an incoming MQTT packet
		msg, err := mqtt.DecodePacket(reader, maxSize)
		if err != nil {
			logger.Error("mqtt decode packet failed", zap.Error(err))
			return
		}

		// Handle the receive
		if err := c.onReceive(msg); err != nil {
			logger.Warn("on receive packet failed", zap.Error(err))
			return
		}
	}
}

func (c *Conn) Unread(unread *message.Unread) error {
	if len(unread.Topics) == 0 {
		return nil
	}

	var err error
	msg := message.New()
	msg.Type = message.MsgUnread
	msg.ID, err = c.broker.msgID.MsgID()
	if err != nil {
		msg.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	body, err := unread.Encode()
	if err != nil {
		logger.Warn("unread encode failed", zap.Error(err))
		return err
	}

	msg.Payload = body

	return c.publishCmd([]*message.Message{msg})
}

// onReceive handles an MQTT receive.
func (c *Conn) onReceive(msg mqtt.Message) error {
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
		var topics []string
		// Subscribe for each subscription
		for _, sub := range packet.Subscriptions {
			// @TODO 订阅合法性验证
			if err := c.onSubscribe(string(sub.Topic)); err != nil {
				ack.Qos = append(ack.Qos, 0x80) // 0x80 indicate subscription failure
				continue
			}
			// Append the QoS
			ack.Qos = append(ack.Qos, sub.Qos)
			topics = append(topics, string(sub.Topic))
		}

		// Acknowledge the subscription
		if _, err := ack.EncodeTo(c.socket); err != nil {
			return err
		}
		unread := message.NewUnread()
		// 是否有未读消息
		for _, topic := range topics {
			insertTime, err := c.broker.cache.GetInt64(topic + message.Topic_Msg_InsertTime)
			if err != nil {
				logger.Warn("get get insert time failed", zap.Error(err))
				continue
			}
			ackTime, err := c.broker.cache.GetInt64(topic + "_" + c.username + message.User_Msg_AckTime)
			if err != nil {
				logger.Warn("get incr failed", zap.Error(err))
				continue
			}
			if insertTime > ackTime {
				unread.Topics[topic] = message.HASE_UNREAD_MSG
			}
		}
		if err := c.Unread(unread); err != nil {
			logger.Warn("notify unread msg failed", zap.Error(err))
			return err
		}
	// We got an attempt to unsubscribe from a channel.
	case mqtt.TypeOfUnsubscribe:
		packet := msg.(*mqtt.Unsubscribe)
		ack := mqtt.Unsuback{MessageID: packet.MessageID}
		// Unsubscribe from each subscription
		for _, sub := range packet.Topics {
			if err := c.onUnsubscribe(string(sub.Topic)); err != nil {
				logger.Warn("onUnsubscribe failed", zap.Error(err))
				return err
			}
		}

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
		// @TODO 推送&拉取合法性验证
		packet := msg.(*mqtt.Publish)
		msgs, err := message.Decode(packet.Payload)
		if err != nil {
			logger.Warn("msg decode failed", zap.Error(err))
			return err
		}

		for _, msg := range msgs {
			if err := c.onPublish(packet, msg); err != nil {
				logger.Warn("onPublish failed", zap.Uint64("cid", c.cid), zap.String("userName", c.username), zap.Error(err))
			}
		}

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
		packet := msg.(*mqtt.Puback)
		msgIDinfo, ok := c.pubIDs.Load(packet.MessageID)
		if ok {
			// ack 时间更新
			c.broker.cache.Set(msgIDinfo.(*msgIDInfo).topic+"_"+c.username+message.User_Msg_AckTime, time.Now().Unix(), 0)
			c.storeMsgID(msgIDinfo.(*msgIDInfo).topic, msgIDinfo.(*msgIDInfo).msgID)
			logger.Info("recv Ack", zap.Uint64("cid", c.cid), zap.String("userName", c.username), zap.String("topic", msgIDinfo.(*msgIDInfo).topic), zap.String("msgID", msgIDinfo.(*msgIDInfo).msgID))
		}
		c.pubIDs.Delete(packet.MessageID)

		break

	default:
		logger.Warn("unknown msg type", zap.Uint8("type", msg.Type()))
		return fmt.Errorf("unknown msg type, %d", msg.Type())
	}
	return nil
}

func (c *Conn) storeMsgID(topic, msgID string) {
	c.topics.Store(topic, msgID)
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
	c.cid = c.broker.uid.Uid()
	return true
}

func (c *Conn) onSubscribe(topic string) error {
	// 建立topic和clientID直接映射关系
	if err := c.broker.subscribe(topic, c.cid, c); err != nil {
		logger.Warn("subscribe failed", zap.String("userName", c.username), zap.String("clientID", c.clientID), zap.Uint64("cid", c.cid), zap.Error(err))
		return err
	}

	// 在conn缓存订阅的topic，在下线的时候用来清除全局的topic中的cid
	c.topics.Store(topic, "")
	return nil
}

func (c *Conn) onUnsubscribe(topic string) error {
	if err := c.broker.cache.Unsubscribe(c.username, topic); err != nil {
		return err
	}

	if err := c.broker.unSubscribe(topic, c.cid); err != nil {
		return err
	}
	return nil
}

func (c *Conn) onPublish(packet *mqtt.Publish, msg *message.Message) error {
	switch msg.Type {
	case message.MsgPub:
		// @TODO 消息存储, 申请ID纪录转换
		msgID, err := gBroker.msgID.MsgID()
		if err != nil {
			logger.Warn("get msg id failed", zap.Error(err))
			return err
		}

		logger.Info("recv publish msg", zap.Uint64("cid", c.cid), zap.String("topic", string(packet.Topic)), zap.String("msgid", msgID), zap.String("originalID", msg.ID))

		if err := c.broker.storage.Store([]*message.Message{msg}, []string{msgID}); err != nil {
			logger.Warn("store msg failed", zap.Error(err))
			return err
		}

		// 检测在线用并推送
		if err := c.broker.publish(c.cid, msgID, msg); err != nil {
			logger.Warn("push online failed", zap.Error(err))
		}
		insertTime := time.Now().Unix()
		// Topic 下插入时间更新
		c.broker.cache.Set(msg.Topic+message.Topic_Msg_InsertTime, insertTime, 0)

		// 用户Ack时间更新
		c.broker.cache.Set(msg.Topic+"_"+c.username+message.User_Msg_AckTime, insertTime, 0)
		// 将消息推送到其他集群上
		if _, err = c.broker.cluster.SyncMsg(msg); err != nil {
			logger.Warn("cluster sysnc msg failed", zap.Error(err))
		}

		break
	// 消息拉取
	case message.MsgPull:
		ackTime := time.Now().Unix()
		c.broker.cache.Set(msg.Topic+"_"+c.username+message.User_Msg_AckTime, ackTime, 0)

		offset, checkTime, err := message.UnPackPullMsg(msg.Payload)
		if err != nil {
			return err
		}
		if offset < 0 || checkTime < 0 {
			return fmt.Errorf("the pull offset %d checkTime %d equal/smaller than 0", offset, checkTime)
		}

		// 未订阅该主题
		if _, ok := c.topics.Load(msg.Topic); !ok {
			return errors.New("pull messages without subscribe the topic:" + msg.Topic)
		}

		msgs, err := c.broker.storage.Get(msg.Topic, offset, checkTime)
		if err != nil {
			logger.Warn("get msg failed", zap.Uint64("cid", c.cid), zap.String("userName", c.username), zap.Error(err))
			return err
		}

		if len(msgs) > 0 {
			c.sendQueue <- msgs
		}
		break
	default:
		return fmt.Errorf("unknow msg type, type is %d", msg.Type)
	}
	return nil
}

// Publish ...
func (c *Conn) Publish(packet *message.Message) error {
	c.sendQueue <- []*message.Message{packet}
	return nil
}

func (c *Conn) publish(msgs []*message.Message) error {
	if len(msgs) == 0 {
		return nil
	}

	packetMsgID := msgs[len(msgs)-1].ID
	var isCompress byte
	if len(msgs) <= 0 {
		return fmt.Errorf("msgs len is zero")
	} else if len(msgs) > 1 {
		isCompress = message.Compress
	} else {
		isCompress = message.NoCompress
	}

	payload, err := message.Encode(msgs, isCompress)
	if err != nil {
		logger.Warn("packet encode failed", zap.Uint64("cid", c.cid), zap.Error(err))
		return err
	}

	mqttID, _ := c.mqttID()
	msg := mqtt.Publish{
		Header: &mqtt.StaticHeader{
			QOS: 1,
		},
		Topic:     []byte(msgs[0].Topic),
		MessageID: mqttID,
		Payload:   payload,
	}
	_, err = msg.EncodeTo(c.socket)
	if err == nil {
		c.pubIDs.Store(mqttID, newMsgIDInfo(msgs[0].Topic, packetMsgID))
	}
	return err
}

func (c *Conn) publishCmd(msgs []*message.Message) error {
	var isCompress byte
	if len(msgs) <= 0 {
		return fmt.Errorf("msgs len is zero")
	} else if len(msgs) > 1 {
		isCompress = message.Compress
	} else {
		isCompress = message.NoCompress
	}

	payload, err := message.Encode(msgs, isCompress)
	if err != nil {
		logger.Warn("packet encode failed", zap.Uint64("cid", c.cid), zap.Error(err))
		return err
	}

	msgID, _ := c.mqttID()
	msg := mqtt.Publish{
		Header: &mqtt.StaticHeader{
			QOS: 1,
		},
		Topic:     []byte(msgs[0].Topic),
		MessageID: msgID,
		Payload:   payload,
	}
	_, err = msg.EncodeTo(c.socket)
	if err != nil {
		logger.Warn("msg encode failed", zap.Uint64("cid", c.cid), zap.Error(err))
		return err
	}
	return nil
}

type msgIDInfo struct {
	topic string
	msgID string
}

func newMsgIDInfo(topic string, msgID string) *msgIDInfo {
	return &msgIDInfo{
		topic: topic,
		msgID: msgID,
	}
}

// Close terminates the connection.
func (c *Conn) Close() error {
	if r := recover(); r != nil {
		//logging.LogAction("closing", fmt.Sprintf("panic recovered: %s \n %s", r, debug.Stack()))
	}
	c.topics.Range(func(topic, msgID interface{}) bool {
		_ = c.broker.logout(topic.(string), c.cid)
		if msgID.(string) != "" {
			_ = c.broker.cache.StoreMsgID(c.username, topic.(string), msgID.(string))
		}
		return true
	})

	return c.socket.Close()
}
