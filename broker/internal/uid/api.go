package uid

import "github.com/congim/xpush/broker/internal/uid/atomic"

type UIDs interface {
	Uid() uint64
}

func New() UIDs {
	return atomic.NewUIDS()
}
