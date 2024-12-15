package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AcUserModel = (*customAcUserModel)(nil)

type (
	// AcUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAcUserModel.
	AcUserModel interface {
		acUserModel
	}

	customAcUserModel struct {
		*defaultAcUserModel
	}
)

// NewAcUserModel returns a model for the database table.
func NewAcUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AcUserModel {
	return &customAcUserModel{
		defaultAcUserModel: newAcUserModel(conn, c, opts...),
	}
}
