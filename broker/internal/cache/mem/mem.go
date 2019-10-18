package mem

import "log"

type Mem struct {
}

func (m *Mem) Init() error {
	return nil
}

func (m *Mem) Logout(cid uint64, name string) error {
	log.Println(cid, name)
	return nil
}

func (m *Mem) Login(uint64, string) error {
	return nil
}

func New() *Mem {
	return &Mem{}
}
