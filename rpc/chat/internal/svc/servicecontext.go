/*
 * @Author: Leeson
 * @Date: 2024-11-23 23:33:38
 */
package svc

import "AiChatPartner/rpc/chat/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
