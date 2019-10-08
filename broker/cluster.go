package broker

import (
	"fmt"

	"github.com/congim/xpush/broker/internal/cluster"
	"go.uber.org/zap"
)

func notify(event *cluster.Event) error {
	switch event.Type {
	case cluster.Join:
		gBroker.logger.Debug("Join", zap.String("name", event.Name))
		break
	case cluster.Leave:
		gBroker.logger.Debug("Leave", zap.String("name", event.Name))
		break
	case cluster.Update:
		gBroker.logger.Debug("Update", zap.String("name", event.Name))
		break
	default:
		return fmt.Errorf("unknow type, type is %d", event.Type)
	}
	return nil
}
