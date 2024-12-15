package logic

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BeginTransactionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBeginTransactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BeginTransactionLogic {
	return &BeginTransactionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 事务管理
func (l *BeginTransactionLogic) BeginTransaction(in *db.BeginTransactionRequest) (*db.BeginTransactionResponse, error) {
	// todo: add your logic here and delete this line

	return &db.BeginTransactionResponse{}, nil
}
