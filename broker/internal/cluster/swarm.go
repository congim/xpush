package cluster

import (
	"io/ioutil"
	"sync"
	"time"

	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

var _ Cluster = (*Swarm)(nil)

type Swarm struct {
	sync.Mutex
	conf   *config.Cluster
	logger *zap.Logger
	notify func(*Event) error
	hosts  *memberlist.Memberlist
	peers  sync.Map
}

func newSwarm(conf *config.Cluster, logger *zap.Logger, notify func(*Event) error) *Swarm {
	return &Swarm{
		conf:   conf,
		logger: logger,
		notify: notify,
	}
}

func (s *Swarm) Start() error {
	go func() {
		if err := startRPC(s.conf.Name, s.logger, s.notify); err != nil {
			s.logger.Fatal("init server failed", zap.Error(err))
			return
		}
		return
	}()

	time.Sleep(3 * time.Second)

	conf := memberlist.DefaultLocalConfig()
	conf.Name = s.conf.Name
	conf.BindPort = s.conf.BindPort
	conf.AdvertisePort = s.conf.BindPort
	conf.BindAddr = s.conf.BindAddr
	conf.Events = newEvent(s, s.logger)
	conf.LogOutput = ioutil.Discard

	hosts, err := memberlist.Create(conf)
	if err != nil {
		s.logger.Warn("create member list failed", zap.Error(err))
		return err
	}

	s.hosts = hosts

	//go func() {
	//	for {
	//		fmt.Println("---------------start----------------")
	//		for _, member := range s.peers.Members() {
	//			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	//		}
	//		fmt.Println("---------------end----------------")
	//		time.Sleep(time.Second * 3)
	//	}
	//}()
	return nil
}

func (s *Swarm) Join() error {
	if _, err := s.hosts.Join(s.conf.Seeds); err != nil {
		s.logger.Warn("peers join failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Swarm) Leave(timeout time.Duration) error {
	if err := s.hosts.Leave(timeout); err != nil {
		s.logger.Warn("peers leave failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Swarm) Close() error {
	if s.hosts != nil {
		_ = s.hosts.Leave(time.Second * 5)
	}
	return nil
}

//func (s *Swarm) OnMessage(peerName string, msg *message.Message) (*message.Reply, error) {
//	peer, ok := s.peers.Load(peerName)
//	if !ok {
//		s.logger.Warn("unfind peer", zap.String("peerName", peerName))
//		return &message.Reply{}, fmt.Errorf("unfind peerName, peer name is %s", peerName)
//	}
//	reply, err := peer.(*Peer).OnMessage(msg)
//	return reply, err
//}

// SyncMsg 同步消息到其他peer
func (s *Swarm) SyncMsg(msg *message.Message) ([]*message.Reply, error) {
	var replys []*message.Reply
	s.peers.Range(func(peerName, peer interface{}) bool {
		reply, err := peer.(*Peer).OnMessage(msg)
		if err != nil {
			s.logger.Warn("on all message failed", zap.String("peerName", peerName.(string)), zap.String("topic", msg.Topic), zap.Error(err))
		}
		replys = append(replys, reply)
		return true
	})

	return replys, nil
}
