package channel

type Channels interface {
}

var _ Channels = (*Channels)(nil)

type channels struct {
	channels map[string]Channel
}

func newChannels() *channels {
	return &channels{
		channels: make(map[string]Channel),
	}
}

func New() Channels {
	return newChannels()
}
