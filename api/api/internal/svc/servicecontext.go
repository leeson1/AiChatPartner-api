/*
 * @Author: Leeson
 * @Date: 2024-11-29 17:51:51
 */
package svc

import (
	"AiChatPartner/api/api/internal/config"
	"AiChatPartner/rpc/chat/chatclient"
	"AiChatPartner/rpc/db/client/databaseservice"
	"AiChatPartner/rpc/db/client/redisservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ChatClient chatclient.Chat
	DbServer   databaseservice.DatabaseService
	RdsServer  redisservice.RedisService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ChatClient: chatclient.NewChat(zrpc.MustNewClient(c.Chat)),
		DbServer:   databaseservice.NewDatabaseService(zrpc.MustNewClient(c.Dbserver)),
		RdsServer:  redisservice.NewRedisService(zrpc.MustNewClient(c.Dbserver)),
	}
}
