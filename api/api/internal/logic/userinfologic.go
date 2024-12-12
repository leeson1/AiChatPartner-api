/*
 * @Author: Leeson
 * @Date: 2024-12-01 17:32:24
 */
package logic

import (
	"context"

	"AiChatPartner/api/api/internal/svc"
	"AiChatPartner/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoRsp, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("[UserInfoLogic] ...")

	resp = &types.UserInfoRsp{
		Name: "admin",
	}

	return resp, nil
}
