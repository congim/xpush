package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/congim/xpush/example/client/basic"
	"github.com/congim/xpush/pkg/message"
	"github.com/congim/xpush/pkg/network/mqtt"
	protobuf "github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
)

type Client struct {
	socket   net.Conn
	userName string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Init(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("conn server failed", err)
		return err
	}
	c.socket = conn
	return nil
}

func (c *Client) Connect(userName string) error {
	c.userName = userName
	connPacket := mqtt.Connect{}
	connPacket.Username = []byte(userName)
	connPacket.UsernameFlag = true
	if _, err := connPacket.EncodeTo(c.socket); err != nil {
		log.Println("mqtt conn failed", err)
		return err
	}
	return nil
}

func (c *Client) loopRead(wg *sync.WaitGroup) {
	defer func() {
		log.Println("loopRead 退出")
		wg.Done()
	}()
	reader := bufio.NewReaderSize(c.socket, 65536)
	for {
		_ = c.socket.SetDeadline(time.Now().Add(time.Second * time.Duration(10)))
		// Decode an incoming MQTT packet
		msg, err := mqtt.DecodePacket(reader, 65536)
		if err != nil {
			log.Println("DecodePacket error", err)
			return
		}

		// Handle the receive
		if err := c.onReceive(msg); err != nil {
			log.Println("onReceive error", err)
			return
		}

	}
}

func (c *Client) ping(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for {
		mqttPing := mqtt.Pingreq{}
		if _, err := mqttPing.EncodeTo(c.socket); err != nil {
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) Subscribe(topics []string) error {
	tQoss := make([]mqtt.TopicQOSTuple, len(topics))
	for index, topic := range topics {
		tQos := mqtt.TopicQOSTuple{
			Topic: []byte(topic),
			Qos:   1,
		}
		tQoss[index] = tQos
	}

	mqttSub := mqtt.Subscribe{Subscriptions: tQoss}
	if _, err := mqttSub.EncodeTo(c.socket); err != nil {
		return err
	}

	return nil
}

// onReceive handles an MQTT receive.
func (c *Client) onReceive(msg mqtt.Message) error {
	switch msg.Type() {
	case mqtt.TypeOfPingresp:
		log.Println("pingresp")
		return nil
	case mqtt.TypeOfConnack:
		log.Println("TypeOfConnack")
		break
	case mqtt.TypeOfDisconnect:
		log.Println("TypeOfDisconnect")
		return nil
	case mqtt.TypeOfSuback:
		log.Println("TypeOfSuback")
		break
	case mqtt.TypeOfPuback:
		//packet := msg.(*mqtt.Pubrec)
		log.Println("TypeOfPuback")
		break
	case mqtt.TypeOfPublish:
		log.Println("TypeOfPublish")
		packet := msg.(*mqtt.Publish)
		// zeus 测试
		// if err := msgDecode(msg.(*mqtt.Publish)); err != nil {
		// 	log.Println("msgDecode", err)
		// 	return err
		// }

		// xpush测试
		msgs, err := message.Decode(packet.Payload)
		if err != nil {
			log.Println("decode failed", err)
		} else {
			for _, msg := range msgs {
				if msg.Type == message.MsgPub {
					log.Print("获得的消息-->>>", msg, string(msg.Payload))
				} else if msg.Type == message.NewMsg {
					unread := message.NewUnread()
					if err := unread.Decode(msg.Payload); err != nil {
						log.Println("unread decode failed", err)
					}
					for topic, isRead := range unread.Topics {
						log.Println("主题", topic, "未读消息标志为", isRead)
						newMsg := message.New()
						newMsg.Topic = topic
						newMsg.Type = message.MsgPull
						newMsg.ID = time.Now().String()
						newMsg.Payload = message.PackPullMsg(10, []byte("10060"))
						body, err := message.Encode([]*message.Message{newMsg}, message.NoCompress)
						if err != nil {
							log.Print("msg encode faileld", err)
							return err
						}

						mqttPub := &mqtt.Publish{
							Header: &mqtt.StaticHeader{
								QOS:    1,
								Retain: false,
								DUP:    false,
							},
							Topic:     []byte(topic),
							Payload:   body,
							MessageID: 1,
						}

						_, err = mqttPub.EncodeTo(c.socket)
						if err != nil {
							log.Println("publish failed", err)
							return err
						}

					}

				}

			}
		}

		if packet.Header.QOS > 0 {
			ack := mqtt.Puback{
				MessageID: packet.MessageID,
			}
			if _, err := ack.EncodeTo(c.socket); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown msg type, %d", msg.Type())
	}
	return nil
}

func msgDecode(msg *mqtt.Publish) error {
	msgsPack := &basic.MsgsPack{}
	if err := protobuf.Unmarshal(msg.Payload, msgsPack); err != nil {
		log.Println("unmarshal failed", err)
		return nil
	}

	log.Println("CompressMsg", *msgsPack.CompressMsg)
	unitPlace := *msgsPack.CompressMsg / 1 % 10
	if unitPlace == basic.ActionCompress {
		cb, err := snappy.Decode(nil, msgsPack.Body)
		if err != nil {
			log.Println("snappy failed", err)
			return err
		}
		msgsPack.Body = cb
	}

	msgs := &basic.Msgs{}
	if err := protobuf.Unmarshal(msgsPack.Body, msgs); err != nil {
		log.Println("unmarshal failed", err)
		return nil
	}
	for _, msg := range msgs.MsgList {
		log.Println("消息列表", string(msg.Body))
	}
	//if *msgsPack.CompressMsg == basic.ActionCompress {
	//
	//}

	return nil
}

func main() {
	log.Println("输入参数为: ", os.Args)
	if len(os.Args) < 4 {
		log.Println("请输入: mqttAddr userName topic")
		return
	}
	client := NewClient()
	if err := client.Init(os.Args[1]); err != nil {
		log.Print("客户端初始化失败", err, os.Args[1])
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go client.loopRead(wg)

	// 链接Conn
	if err := client.Connect(os.Args[2]); err != nil {
		log.Print("客户端订阅主题失败", err, os.Args[1])
		return
	}

	go client.ping(wg)

	// 订阅
	if err := client.Subscribe([]string{os.Args[3]}); err != nil {
		log.Print("客户端订阅主题失败", err, os.Args[1])
		return
	}

	for {
		time.Sleep(5 * time.Second)
		if err := client.push(os.Args[3]); err != nil {
			log.Println("客户端推送消息失败", err, os.Args[1])
			return
		}
	}
	wg.Wait()
}

var gMsgID int32

func getMsgID() int32 {
	return 10000 + atomic.AddInt32(&gMsgID, 1)
}

func (c *Client) push(topic string) error {
	var msgs []*message.Message
	for index := 0; index < 10; index++ {
		msg := message.New()
		msg.Type = message.MsgPub
		msg.Topic = topic
		msg.ID = fmt.Sprintf("%s-%d", c.userName, getMsgID())
		msg.Payload = []byte("hello xpush !")
		msgs = append(msgs, msg)
	}
	payload, err := message.Encode(msgs, message.NoCompress)
	if err != nil {
		log.Println("message encode failed", err)
		return err
	}

	mqttPub := &mqtt.Publish{
		Header: &mqtt.StaticHeader{
			QOS:    1,
			Retain: false,
			DUP:    false,
		},
		Topic:     []byte(topic),
		Payload:   payload,
		MessageID: 1,
	}

	_, err = mqttPub.EncodeTo(c.socket)
	if err != nil {
		log.Println("publish failed", err)
		return err
	}
	return nil
}
