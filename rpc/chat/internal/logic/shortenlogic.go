package logic

import (
	"context"

	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/chat/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *chat.ShortenReq) (*chat.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &chat.ShortenResp{}, nil
}
