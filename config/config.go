package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Common   *Common
	Broker   *Broker
	Listener *Listener
	Limit    *LimitConfig
	Cluster  *Cluster
	Storage  *Storage
	Cache    *Cache
}

type Cluster struct {
	Name     string
	BindAddr string
	BindPort int
	Seeds    []string
}

type LimitConfig struct {
	MessageSize int64
}

type Common struct {
	Version  string
	LogLevel string
}

type Broker struct {
}

type Listener struct {
	ReadTimeOut uint16
	ListenAddr  string
	IsTLS       bool
	Certificate string
	PrivateKey  string
}

type Storage struct {
	Name    string
	Threads int
	DBSpace string
}

type Cache struct {
	Name  string
	Redis *Redis
}

type Redis struct {
	IsCluster bool
	Addr      string
	Addrs     []string
}

// New config struct
func New(path string) (*Config, error) {
	conf := &Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read file failed, err msg is", err)
		return nil, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Println("unmarshal failed, err msg is", err)
		return nil, err
	}
	return conf, nil
}
