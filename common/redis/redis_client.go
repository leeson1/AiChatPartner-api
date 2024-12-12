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

func (rs *RedisClient) setExpire(key string, expire int) error {
	if expire > 0 {
		err := rs.Redis.Expire(key, expire)
		if err != nil {
			return err
		}
	}
	return nil
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
	return nil
}

func (rs *RedisClient) Expire(key string, exp int) error {
	err := rs.Redis.Expire(key, exp)
	if err != nil {
		return fmt.Errorf("[redis_client] error to expire key: %s error: %s", key, err)
	}
	return nil
}

func (rs *RedisClient) Set(key, value string) error {
	err := rs.Redis.Set(key, value)
	if err != nil {
		return fmt.Errorf("[redis_client] failed to set key %s: %w", key, err)
	}
	return nil
}

func (rs *RedisClient) Hset(key, field, value string) error {
	err := rs.Redis.Hset(key, field, value)
	if err != nil {
		return fmt.Errorf("[redis_client] error to hset key %s: %w", key, err)
	}
	return nil
}

func (rs *RedisClient) Hmset(key string, fieldsAndValues map[string]string) error {
	err := rs.Redis.Hmset(key, fieldsAndValues)
	if err != nil {
		return fmt.Errorf("[redis_client] error to hmset key %s: %w", key, err)
	}
	return nil
}

func (rs *RedisClient) SetWithExpire(key, value string, expire int) error {
	err := rs.Set(key, value)
	if err != nil {
		return err
	}
	return rs.setExpire(key, expire)
}

func (rs *RedisClient) HsetWithExpire(key, field, value string, expire int) error {
	err := rs.Hset(key, field, value)
	if err != nil {
		return err
	}

	return rs.setExpire(key, expire)
}

func (rs *RedisClient) HmsetWithExpire(key string, fieldsAndValues map[string]string, expire int) error {
	err := rs.Hmset(key, fieldsAndValues)
	if err != nil {
		return err
	}

	return rs.setExpire(key, expire)
}
