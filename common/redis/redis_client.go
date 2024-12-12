/*
 * @Author: LEESON
 * @Date: 2024-12-11 23:38:25
 */
package redis

import (
	"AiChatPartner/common/config"
	"fmt"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisClient struct {
	Redis *redis.Redis
}

var (
	rdsInstance *RedisClient = &RedisClient{}
	rc          redis.RedisConf
	once        sync.Once
)

func InitRedis(c *config.Config) error {
	rc.Host = c.RedisConf.Host
	rc.Pass = c.RedisConf.Pass
	rc.Tls = c.RedisConf.Tls
	rc.Type = c.RedisConf.Type
	logx.Info("host:", rc.Host)
	//...
	return nil
}

func GetRedisClient() *RedisClient {
	once.Do(func() {
		var err error
		rdsInstance.Redis, err = redis.NewRedis(rc)
		if err != nil {
			logx.Errorf("[redis_client::GetRedisClient] failed to create redis client: %v", err)
		}
	})
	return rdsInstance
}

func NewRedisClient(redisConf redis.RedisConf) (*RedisClient, error) {
	rds, err := redis.NewRedis(redisConf)
	if err != nil {
		return nil, err
	}
	rdsInstance = &RedisClient{Redis: rds}
	return &RedisClient{Redis: rds}, nil
}

func (rs *RedisClient) Set(key string, value string, expire int) error {
	var err error
	err = rs.Redis.Set(key, value)
	if err != nil {
		return fmt.Errorf("[redis_client] failed to set key %s: %w", key, err)
	}
	if expire > 0 {
		err = rs.Redis.Expire(key, expire)
		if err != nil {
			return fmt.Errorf("[redis_client] failed to expire key %s: %w", key, err)
		}
	}
	return err
}

func (rs *RedisClient) Get(key string) (string, error) {
	result, err := rs.Redis.Get(key)
	if err != nil {
		return "", fmt.Errorf("[redis_client] failed to get key %s: %w", key, err)
	}
	return result, nil
}

func (rs *RedisClient) Del(key string) error {
	_, err := rs.Redis.Del(key)
	if err != nil {
		return fmt.Errorf("[redis_client] failed to delete key %s: %w", key, err)
	}
	return err
}

func (rs *RedisClient) Expire(key string, exp int) error {
	err := rs.Redis.Expire(key, exp)
	if err != nil {
		return fmt.Errorf("[redis_client] set redis error: %s", err)
	}
	return nil
}
