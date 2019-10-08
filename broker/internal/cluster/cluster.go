package cluster

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/congim/xpush/config"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

var _ Cluster = (*cluster)(nil)

type cluster struct {
	conf   *config.Cluster
	logger *zap.Logger
	notify func(*Event) error
	peers  *memberlist.Memberlist
}

func newCluster(conf *config.Cluster, logger *zap.Logger, notify func(*Event) error) *cluster {
	return &cluster{
		conf:   conf,
		logger: logger,
		notify: notify,
	}
}

func (c *cluster) Start() error {
	conf := memberlist.DefaultLocalConfig()
	hostname, _ := os.Hostname()
	conf.Name = hostname + "-" + fmt.Sprintf("%d", c.conf.BindPort)
	//conf.Name = c.conf.Name
	conf.BindPort = c.conf.BindPort
	conf.AdvertisePort = c.conf.BindPort
	conf.BindAddr = c.conf.BindAddr
	conf.Events = newEvent(c, c.logger)
	conf.LogOutput = ioutil.Discard

	peers, err := memberlist.Create(conf)
	if err != nil {
		c.logger.Warn("create member list failed", zap.Error(err))
		return err
	}

	c.peers = peers
	//go func() {
	//	for {
	//		fmt.Println("---------------start----------------")
	//		for _, member := range c.peers.Members() {
	//			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	//		}
	//		fmt.Println("---------------end----------------")
	//		time.Sleep(time.Second * 3)
	//	}
	//}()
	return nil
}

func (c *cluster) Join() error {
	if _, err := c.peers.Join(c.conf.Seeds); err != nil {
		c.logger.Warn("peers join failed", zap.Error(err))
		return err
	}
	return nil
}

func (c *cluster) Leave(timeout time.Duration) error {
	if err := c.peers.Leave(timeout); err != nil {
		c.logger.Warn("peers leave failed", zap.Error(err))
		return err
	}
	return nil
}

func (c *cluster) Close() error {
	if c.peers != nil {
		_ = c.peers.Leave(time.Second * 5)
	}
	return nil
}
