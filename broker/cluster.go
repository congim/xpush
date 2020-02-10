package broker

import (
	"fmt"

	"github.com/congim/xpush/broker/internal/cluster"
	"go.uber.org/zap"
)

func notify(event *cluster.Event) error {
	switch event.Type {
	case cluster.Join:
		logger.Info("Join", zap.String("name", event.Name))
		break
	case cluster.Leave:
		logger.Info("Leave", zap.String("name", event.Name))
		break
	case cluster.Update:
		logger.Info("Update", zap.String("name", event.Name))
		break
	case cluster.Pub:
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
