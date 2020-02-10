package broker

import (
	"fmt"
	"log"

	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

func notify(event *message.Event) error {
	switch event.Type {
	case message.ClusterJoin:
		logger.Info("Join", zap.String("name", event.Name))
		break
	case message.ClusterLeave:
		logger.Info("Leave", zap.String("name", event.Name))
		break
	case message.ClusterUpdate:
		logger.Info("Update", zap.String("name", event.Name))
		break
	case message.MsgPub:
		log.Println("收到其他节点到消息")
		for _, msg := range event.Msgs {
			if err := gBroker.notify(msg); err != nil {
				logger.Warn("publish msg failed", zap.String("topic", msg.Topic), zap.String("msgID", msg.ID), zap.Error(err))
				continue
			}
		}
	default:
		return fmt.Errorf("unknow type, type is %d", event.Type)
	}
	return nil
}
