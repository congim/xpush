package topic

import "github.com/congim/xpush/pkg/message"

type Topic interface {
	Subscribe(string) error
	UnSubscribe(string) error
	OnMessage(message *message.Message) error
}
