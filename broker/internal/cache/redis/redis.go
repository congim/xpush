package redis

import (
	"log"

	"github.com/congim/xpush/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
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

// Inc .
func (r *Redis) Inc(topic string, msgID string) error {
	statusCmd := r.client.Set(topic, msgID, 0)
	if statusCmd.Err() != nil {
		r.logger.Warn("cache Publish failed", zap.Error(statusCmd.Err()))
		return statusCmd.Err()
	}
	return nil
}

// Unread  ..
func (r *Redis) Unread(topic string, userName string) (bool, error) {
	log.Println(topic, userName)
	newMsgInfo := r.client.Get(topic)
	newMsgID, err := newMsgInfo.Result()
	log.Println(userName, topic, newMsgID)
	if err == redis.Nil {
		return false, nil
	}

	lastMsgInfo := r.client.HGet(LAST_MSG_ID+topic, userName)
	lastMsgID, err := lastMsgInfo.Result()
	log.Println(userName, LAST_MSG_ID+topic, "lastMsgID", lastMsgID)

	if err == redis.Nil {
		return true, nil
	}

	if newMsgID != lastMsgID {
		return true, nil
	}

	return false, nil
}

// New new reids
func New(conf *config.Redis, l *zap.Logger) *Redis {
	return &Redis{
		conf:   conf,
		logger: l,
	}
}
