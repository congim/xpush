package redis

import (
	"github.com/congim/xpush/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"log"
)

const (
	LAST_MSG_ID string = "LASTMSGID-"
)

// Redis ....
type Redis struct {
	logger *zap.Logger
	conf   *config.Redis
	client *redis.Client
	// clusterClient *redis.ClusterClient
}

// Init ...
func (r *Redis) Init() error {
	// if !r.conf.IsCluster {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.conf.Addr,     // use default Addr
		Password: r.conf.Password, // no password set
		DB:       0,               // use default DB
	})
	status := rdb.Ping()
	if status.Err() != nil {
		r.logger.Error("redis ping failed", zap.Error(status.Err()))
		return status.Err()
	}
	r.client = rdb
	// } else {
	// 	rdb := redis.NewClusterClient(&redis.ClusterOptions{
	// 		Addrs: r.conf.Addrs,
	// 	})
	// 	status := rdb.Ping()
	// 	if status.Err() != nil {
	// 		r.logger.Error("redis ping failed", zap.Error(status.Err()))
	// 		return status.Err()
	// 	}
	// 	r.clusterClient = rdb
	// }
	return nil
}

// Unsubscribe ...
func (r *Redis) Unsubscribe(userName string, topic string) error {
	var intCmd *redis.IntCmd
	intCmd = r.client.HDel(LAST_MSG_ID+topic, userName)
	if intCmd.Err() != nil {
		r.logger.Warn("cache Unsubscribe failed", zap.Error(intCmd.Err()), zap.String("userName", userName), zap.String("topic", topic))
		return intCmd.Err()
	}
	return nil
}

// StoreMsgID .
func (r *Redis) StoreMsgID(userName string, topic string, msgID string) error {
	statusCmd := r.client.HSet(LAST_MSG_ID+topic, userName, msgID)
	if statusCmd.Err() != nil {
		r.logger.Warn("cache StoreMsgID failed", zap.Error(statusCmd.Err()))
		return statusCmd.Err()
	}
	return nil
}

// Publish .
func (r *Redis) Publish(topic string, msgID string) error {
	statusCmd := r.client.Set(topic, msgID, 0)
	if statusCmd.Err() != nil {
		r.logger.Warn("cache Publish failed", zap.Error(statusCmd.Err()))
		return statusCmd.Err()
	}
	return nil
}

// Unread  ..
func (r *Redis) Unread(topic string, userName string) (bool, error) {
	newMsgInfo := r.client.Get(topic)
	newMsgID, err := newMsgInfo.Result()
	if err == redis.Nil {
		log.Println("---1")
		return false, nil
	}

	lastMsgInfo := r.client.HGet(LAST_MSG_ID+topic, userName)
	lastMsgID, err := lastMsgInfo.Result()
	if err == redis.Nil {
		log.Println("---2")
		return true, nil
	}

	if newMsgID != lastMsgID {
		log.Println("---3")
		return true, nil
	}
	log.Println("---4")
	return false, nil
}

// New new reids
func New(conf *config.Redis, l *zap.Logger) *Redis {
	return &Redis{
		conf:   conf,
		logger: l,
	}
}

// // Logout ...
// func (r *Redis) Logout(topic string, msgID string) error {
// 	return nil
// }

// // Login ....
// func (r *Redis) Login(userName string, brokerName string) error {
// 	return nil
// }

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

// // UnRead ..
// func (r *Redis) UnRead(userName string, topics []string) (*message.UnRead, error) {
// 	unRead := message.NewUnRead()
// 	for _, topic := range topics {
// 		var totalCmd *redis.StringCmd
// 		//	total
// 		if !r.conf.IsCluster {
// 			totalCmd = r.client.Get(topic)
// 		} else {
// 			totalCmd = r.client.Get(topic)
// 		}
// 		_, err := totalCmd.Result()
// 		if err == redis.Nil {
// 			unRead.Topics[topic] = 0
// 			continue
// 		}

// 		// received
// 		var recvCmd *redis.StringCmd
// 		recvCmd = r.client.HGet(userName, topic)
// 		if recvCmd.Err() != nil {
// 			r.logger.Warn("HGet failed", zap.Error(recvCmd.Err()))
// 			unRead.Topics[topic] = 0
// 			continue
// 		}
// 		total, err := totalCmd.Int64()
// 		if err != nil {
// 			unRead.Topics[topic] = 0
// 			continue
// 		}
// 		recv, err := recvCmd.Int64()
// 		if err != nil {
// 			unRead.Topics[topic] = 0
// 			continue
// 		}
// 		if total >= recv {
// 			unRead.Topics[topic] = total - recv
// 		} else {
// 			unRead.Topics[topic] = 0
// 		}

// 		log.Println("主题", topic, "未接收的消息量为", total-recv, ",total=", total, ",recv=", recv)
// 	}
// 	return unRead, nil
// }

// PubCount ...
// func (r *Redis) PubCount(userName string, topic string, count int) error {
// 	r.client.IncrBy(topic, int64(count))
// 	r.client.HIncrBy(userName, topic, int64(count))
// 	return nil
// }

// // Ack ...
// func (r *Redis) Ack(userName string, topic string, count uint64) error {
// 	var intCmd *redis.IntCmd
// 	intCmd = r.client.HIncrBy(userName, topic, int64(count))
// 	if intCmd.Err() != nil {
// 		r.logger.Warn("HIncrBy failed", zap.Error(intCmd.Err()))
// 		return intCmd.Err()
// 	}
// 	return nil
// }

// Subscribe ...
// func (r *Redis) Subscribe(userName string, topic string) error {
// // 保存用户订阅哪些topic
// var intCmd *redis.IntCmd
// // if !r.conf.IsCluster {
// intCmd = r.client.HIncrBy(userName, topic, 0)
// // } else {
// // intCmd = r.clusterClient.HIncrBy(userName, topic, 0)
// // }

// if intCmd.Err() != nil {
// 	r.logger.Warn("SAdd failed", zap.Error(intCmd.Err()))
// 	return intCmd.Err()
// }

// recvCount, err := intCmd.Result()
// if err != nil {
// 	r.logger.Warn("intCmd result error", zap.Error(err))
// 	return intCmd.Err()
// }

// // 已接收消息条数如果为0的情况，那么代表第一次订阅，需要将这个主题的发送量和用户已接收量保持一致，这样不会造成历史消息为未读情况
// if recvCount == 0 {
// 	var strCmd *redis.StringCmd
// 	// 获取已读消息
// 	if !r.conf.IsCluster {
// 		strCmd = r.client.Get(topic)
// 	} else {
// 		strCmd = r.client.Get(topic)
// 	}

// 	_, err := strCmd.Result()
// 	if err == redis.Nil {
// 		// if !r.conf.IsCluster {
// 		r.client.Set(topic, 1, 0)
// 		// } else {
// 		// 	r.clusterClient.Set(topic, 1, 0)
// 		// }
// 	} else if err != nil {
// 		r.logger.Warn("Get failed", zap.String("topic", topic), zap.Error(strCmd.Err()))
// 		return err
// 	}

// 	// 如果主题消息数为0的情况，那么给这个主题设置发送量为0
// 	sendCount, _ := strCmd.Uint64()
// 	if sendCount == 0 {
// 		sendCount = 1
// 	}
// 	intCmd = r.client.HIncrBy(userName, topic, int64(sendCount))
// 	if intCmd.Err() != nil {
// 		r.logger.Warn("HIncrBy failed", zap.Error(intCmd.Err()))
// 		return intCmd.Err()
// 	}
// }
// 	return nil
// }
