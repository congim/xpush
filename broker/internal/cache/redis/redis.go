package redis

import "log"

type Redis struct {
}

func (r *Redis) Init() error {
	return nil
}

func (r *Redis) Logout(cid uint64) error {
	log.Println(cid)
	return nil
}

func (r *Redis) Login(cid uint64, name string) error {
	log.Println(cid, name)
	return nil
}

func (r *Redis) Get(uint64) (string, bool) {
	return "", false
}

func (r *Redis) Subscribe(userName string, topic string) error {
	return nil
}

func (r *Redis) Unsubscribe(userName string, topic string) error {
	return nil
}

func (r *Redis) PubCount(topic string, count int) error {
	return nil
}

func (r *Redis) Ack(userName string, topic string, count uint64) error {
	return nil
}

func New() *Redis {
	return &Redis{}
}
