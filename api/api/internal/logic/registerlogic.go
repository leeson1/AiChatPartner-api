/*
 * @Author: Leeson
 * @Date: 2024-12-15 14:33:31
 */
package logic

import (
	"context"
	"fmt"

	"AiChatPartner/api/api/internal/svc"
	"AiChatPartner/api/api/internal/types"
	"AiChatPartner/rpc/chat/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRsp, err error) {

	rsp, err := l.svcCtx.ChatClient.Register(l.ctx, &chat.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Error("[RegisterLogic] rpc.Register error: ", err)
		return nil, err
	}

	if rsp.RetCode != 0 {
		return &types.RegisterRsp{RetCode: 1}, fmt.Errorf("username:%s already exists", req.Username)
	}

	logx.Infof("[LoginLogic] register success. username: %s", req.Username)

	return &types.RegisterRsp{RetCode: 0}, nil
}
