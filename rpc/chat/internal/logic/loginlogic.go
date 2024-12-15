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
	"AiChatPartner/rpc/db/db"

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

	dbrsp, err := l.svcCtx.DbServer.Read(l.ctx, &db.ReadRequest{
		TableName: "ac_user",
		Key:       in.Username,
		KeyType:   2,
	})
	// user, err := l.svcCtx.Model.GetUserByUsername(l.ctx, in.Username)
	// user, err := mysql.GetMysqlClient().GetUserByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, fmt.Errorf("[rpc/chat Login] get user by username:[%s] error: %s", in.Username, err)
	}

	pass := dbrsp.Data["password"]
	if in.Password != pass {
		return nil, fmt.Errorf("[rpc/chat Login] user:[%s] password error. ", in.Username)
	}

	logx.Info("[rpc/chat Login] login success. user: ", in.Username)

	return &chat.LoginRsp{}, nil
}
