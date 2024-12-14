/*
 * @Author: Leeson
 * @Date: 2024-12-14 17:25:47
 */
package logic

import (
	"context"
	"fmt"

	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/chat/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *chat.LoginReq) (*chat.LoginRsp, error) {

	// var user *model.AcUser
	// var err error
	// user, err := l.svcCtx.Model.FindOne(l.ctx, 1001)
	// if err != nil {
	// 	return nil, fmt.Errorf("[rpc/chat Login] mysql error: %s", err)
	// }

	user, err := l.svcCtx.Model.GetUserByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, fmt.Errorf("[rpc/chat Login] get user by username:[%s] error: %s", in.Username, err)
	}

	if in.Password != user.Password {
		return nil, fmt.Errorf("[rpc/chat Login] user:[%s] password error. ", in.Username)
	}

	logx.Info("[rpc/chat Login] login success. user: ", user.Username)

	return &chat.LoginRsp{}, nil
}
