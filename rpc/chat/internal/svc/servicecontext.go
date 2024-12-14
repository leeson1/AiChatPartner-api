/*
 * @Author: Leeson
 * @Date: 2024-11-23 23:33:38
 */
package svc

import (
	"AiChatPartner/rpc/chat/internal/config"
	"AiChatPartner/rpc/chat/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.AcUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewAcUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
