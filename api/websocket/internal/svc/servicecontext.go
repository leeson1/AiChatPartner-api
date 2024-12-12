/*
 * @Author: Leeson
 * @Date: 2024-11-29 17:52:01
 */
package svc

import (
	"AiChatPartner/api/websocket/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
