package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AcUserModel = (*customAcUserModel)(nil)

type (
	// AcUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAcUserModel.
	AcUserModel interface {
		acUserModel
		withSession(session sqlx.Session) AcUserModel
	}

	customAcUserModel struct {
		*defaultAcUserModel
	}
)

// NewAcUserModel returns a model for the database table.
func NewAcUserModel(conn sqlx.SqlConn) AcUserModel {
	return &customAcUserModel{
		defaultAcUserModel: newAcUserModel(conn),
	}
}

func (m *customAcUserModel) withSession(session sqlx.Session) AcUserModel {
	return NewAcUserModel(sqlx.NewSqlConnFromSession(session))
}
