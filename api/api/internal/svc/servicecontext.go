/*
 * @Author: Leeson
 * @Date: 2024-11-29 17:51:51
 */
package svc

import (
	"AiChatPartner/api/api/internal/config"
	"AiChatPartner/rpc/chat/chatclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ChatClient chatclient.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ChatClient: chatclient.NewChat(zrpc.MustNewClient(c.Chat)),
	}
}
