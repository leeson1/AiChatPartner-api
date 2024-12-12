/*
 * @Author: Leeson
 * @Date: 2024-11-29 17:51:51
 */
package svc

import (
	"AiChatPartner/api/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// Redis:  redis.MustNewRedis(c.RedisConf),
	}
}
