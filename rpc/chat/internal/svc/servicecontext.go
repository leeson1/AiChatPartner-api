/*
 * @Author: Leeson
 * @Date: 2024-11-23 23:33:38
 */
package svc

import (
	"AiChatPartner/rpc/chat/internal/config"
	"AiChatPartner/rpc/db/databaseservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	DbServer databaseservice.DatabaseService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		DbServer: databaseservice.NewDatabaseService(zrpc.MustNewClient(c.Dbserver)),
	}
}
