package redis

import (
	"log"

	"github.com/congim/xpush/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

//
//const (
//	Broker_Prefix string = "broker_"
//)

type Redis struct {
	logger        *zap.Logger
	conf          *config.Redis
	clusterClient *redis.ClusterClient
	client        *redis.Client
}

func (r *Redis) Init() error {
	if !r.conf.IsCluster {
		rdb := redis.NewClient(&redis.Options{
			Addr:     r.conf.Addr, // use default Addr
			Password: "",          // no password set
			DB:       0,           // use default DB
		})
		status := rdb.Ping()
		if status.Err() != nil {
			r.logger.Error("redis ping failed", zap.Error(status.Err()))
			return status.Err()
		}
		r.client = rdb
	} else {
		rdb := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: r.conf.Addrs,
		})
		status := rdb.Ping()
		if status.Err() != nil {
			r.logger.Error("redis ping failed", zap.Error(status.Err()))
			return status.Err()
		}
		r.clusterClient = rdb
	}
	return nil
}

func (r *Redis) Logout(topic string) error {

	return nil
}

func (r *Redis) Login(cid uint64, name string) error {
	log.Println(cid, name)
	return nil
}

func (r *Redis) GetBroker(uint64) (string, bool) {
	return "", false
}

func (r *Redis) Subscribe(userName string, topic string) error {
	// 保存用户订阅哪些topic
	var intCmd *redis.IntCmd
	if !r.conf.IsCluster {
		intCmd = r.client.HIncrBy(userName, topic, 0)
	} else {
		intCmd = r.clusterClient.HIncrBy(userName, topic, 0)
	}

	if intCmd.Err() != nil {
		r.logger.Warn("SAdd failed", zap.Error(intCmd.Err()))
		return intCmd.Err()
	}

	recvCount, err := intCmd.Result()
	if err != nil {
		r.logger.Warn("intCmd result error", zap.Error(err))
		return intCmd.Err()
	}

	// 已接收消息条数如果为0的情况，那么代表第一次订阅，需要将这个主题的发送量和用户已接收量保持一致，这样不会造成历史消息为未读情况
	if recvCount == 0 {
		var strCmd *redis.StringCmd
		// 获取已读消息
		if !r.conf.IsCluster {
			strCmd = r.client.Get(topic)
		} else {
			strCmd = r.client.Get(topic)
		}

		_, err := strCmd.Result()
		if err == redis.Nil {
			if !r.conf.IsCluster {
				r.client.Set(topic, 0, 0)
			} else {
				r.clusterClient.Set(topic, 0, 0)
			}
			return nil
		} else if err != nil {
			r.logger.Warn("Get failed", zap.String("topic", topic), zap.Error(strCmd.Err()))
			return err
		}

		// 如果主题消息数为0的情况，那么给这个主题设置发送量为0
		sendCount, err := strCmd.Uint64()
		if err != nil {
			r.logger.Warn("strCmd Uint64 failed", zap.String("topic", topic), zap.Error(strCmd.Err()))
			return err
		}

		// 保持一致
		if !r.conf.IsCluster {
			intCmd = r.client.HIncrBy(userName, topic, int64(sendCount))
		} else {
			intCmd = r.clusterClient.HIncrBy(userName, topic, int64(sendCount))
		}
		if intCmd.Err() != nil {
			r.logger.Warn("HIncrBy failed", zap.Error(intCmd.Err()))
			return intCmd.Err()
		}
	}
	return nil
}

func (r *Redis) Unsubscribe(userName string, topic string) error {
	var intCmd *redis.IntCmd
	if !r.conf.IsCluster {
		intCmd = r.client.HDel(userName, topic)
	} else {
		intCmd = r.clusterClient.HDel(userName, topic)
	}
	if intCmd.Err() != nil {
		r.logger.Warn("HDel failed", zap.Error(intCmd.Err()), zap.String("userName", userName), zap.String("topic", topic))
		return intCmd.Err()
	}
	return nil
}

func (r *Redis) PubCount(topic string, count int) error {
	return nil
}

func (r *Redis) Ack(userName string, topic string, count uint64) error {
	return nil
}

func New(conf *config.Redis, l *zap.Logger) *Redis {
	return &Redis{
		conf:   conf,
		logger: l,
	}
}

// HSETNX
// HINCRBY
// HGET
//
//func newClusterClient() {
//	rdb := redis.NewClusterClient(&redis.ClusterOptions{
//		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
//	})
//
//	rdb.Ping()
//}
//
//func newClient() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379", // use default Addr
//		Password: "",               // no password set
//		DB:       0,                // use default DB
//	})
//	pong, err := rdb.Ping().Result()
//	fmt.Println(pong, err)
//}
