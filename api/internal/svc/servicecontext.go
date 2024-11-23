package svc

import (
	"AiChatPartner/api/internal/config"
	"AiChatPartner/rpc/chat/chatclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	Chatclient chatclient.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Chatclient: chatclient.NewChat(zrpc.MustNewClient(c.Chat)),
	}
}
