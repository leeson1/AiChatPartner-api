/*
 * @Author: Leeson
 * @Date: 2024-12-15 15:01:46
 */
package databaseservicelogic

import (
	"context"
	"strconv"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"
	"AiChatPartner/rpc/db/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CRUD 操作
func (l *CreateLogic) Create(in *db.CreateRequest) (*db.CreateResponse, error) {

	//string 转 int64
	role, err := strconv.ParseInt(in.Data["role"], 10, 64)
	if err != nil {
		return &db.CreateResponse{Success: false}, err
	}

	l.svcCtx.Model.Insert(l.ctx, &model.AcUser{
		Role:     role,
		Username: in.Data["username"],
		Password: in.Data["password"],
	})

	return &db.CreateResponse{Success: false}, nil
}
