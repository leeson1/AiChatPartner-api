/*
 * @Author: Leeson
 * @Date: 2024-12-15 15:01:46
 */
package databaseservicelogic

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 连接管理
func (l *ConnectLogic) Connect(in *db.ConnectRequest) (*db.ConnectResponse, error) {

	return &db.ConnectResponse{}, nil
}
