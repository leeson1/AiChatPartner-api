/*
 * @Author: Leeson
 * @Date: 2024-11-24 00:18:51
 */
package logic

import (
	"context"

	"AiChatPartner/api/internal/svc"
	"AiChatPartner/api/internal/types"
	"AiChatPartner/rpc/chat/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (resp *types.ExpandResp, err error) {
	// todo: add your logic here and delete this line

	rpcResp, err := l.svcCtx.Chatclient.Expand(l.ctx, &chat.ExpandReq{
		Shorten: req.Shorten,
	})
	if err != nil {
		return nil, err
	}

	return &types.ExpandResp{
		Url: rpcResp.Url,
	}, nil

}
