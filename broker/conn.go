package broker

import (
	"bufio"
	"log"
	"net"
	"time"

	"github.com/congim/xpush/pkg/network/mqtt"
)

type Conn struct {
	broker    *Broker
	socket    net.Conn // The transport used to read and write messages.
	username  string
	password  string
	keepalive uint16
	clientID  string
}

func newConn(conn net.Conn, broker *Broker, readTout uint16) *Conn {
	return &Conn{
		socket:    conn,
		broker:    broker,
		keepalive: readTout,
	}
}

// Process processes the messages.
func (c *Conn) Process() {
	defer func() {
		_ = c.Close()
	}()
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
			//if err := c.onSubscribe(sub.Topic); err != nil {
			//	ack.Qos = append(ack.Qos, 0x80) // 0x80 indicate subscription failure
			//	c.notifyError(err, packet.MessageID)
			//	continue
			//}

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
		return nil

	case mqtt.TypeOfPublish:
		packet := msg.(*mqtt.Publish)
		log.Println("push msg", string(packet.Payload))
		//if err := c.onPublish(packet); err != nil {
		//	logging.LogError("conn", "publish received", err)
		//	c.notifyError(err, packet.MessageID)
		//}
		// Acknowledge the publication
		if packet.Header.QOS > 0 {
			ack := mqtt.Puback{MessageID: packet.MessageID}
			if _, err := ack.EncodeTo(c.socket); err != nil {
				return err
			}
		}
	}

	return nil
}

// onConnect handles the connection authorization
func (c *Conn) onConnect(packet *mqtt.Connect) bool {

	// @TODO 账号密码校验
	c.username = string(packet.Username)
	c.password = string(packet.Password)
	c.clientID = string(packet.ClientID)
	if c.keepalive < packet.KeepAlive {
		c.keepalive = packet.KeepAlive
	}
	return true
}
