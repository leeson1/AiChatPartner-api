/*
 * @Author: Leeson
 * @Date: 2024-12-15 12:31:10
 */
package svc

import (
	"AiChatPartner/rpc/db/internal/config"
	"AiChatPartner/rpc/db/model"

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
