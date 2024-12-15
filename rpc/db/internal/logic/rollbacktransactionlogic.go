package logic

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackTransactionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackTransactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackTransactionLogic {
	return &RollbackTransactionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackTransactionLogic) RollbackTransaction(in *db.RollbackTransactionRequest) (*db.RollbackTransactionResponse, error) {
	// todo: add your logic here and delete this line

	return &db.RollbackTransactionResponse{}, nil
}
