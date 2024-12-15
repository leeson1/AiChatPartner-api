package databaseservicelogic

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommitTransactionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommitTransactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommitTransactionLogic {
	return &CommitTransactionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommitTransactionLogic) CommitTransaction(in *db.CommitTransactionRequest) (*db.CommitTransactionResponse, error) {
	// todo: add your logic here and delete this line

	return &db.CommitTransactionResponse{}, nil
}
