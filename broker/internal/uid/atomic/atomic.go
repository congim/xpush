package atomic

import "sync/atomic"

type UIDS struct {
	uid uint64
}

func (u *UIDS) Uid() uint64 {
	id := atomic.AddUint64(&u.uid, 1)
	return id
}

func NewUIDS() *UIDS {
	return &UIDS{}
}
