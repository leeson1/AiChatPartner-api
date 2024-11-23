package logic

import (
	"context"

	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/chat/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *chat.ExpandReq) (*chat.ExpandResp, error) {
	// todo: add your logic here and delete this line

	return &chat.ExpandResp{}, nil
}
