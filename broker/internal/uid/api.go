package uid

import "github.com/congim/xpush/broker/internal/uid/satomic"

type UIDs interface {
	Uid() uint64
}

func New() UIDs {
	return satomic.NewUIDS()
}
