/*
 * @Author: Leeson
 * @Date: 2024-12-01 17:06:26
 */
package logic

import (
	"context"

	"github.com/zeromicro/x/errors"

	"AiChatPartner/api/api/internal/svc"
	"AiChatPartner/api/api/internal/types"
	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/db/db"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func getJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {

	//生成token
	now := jwt.TimeFunc().Unix()
	token, err := getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, req.Username)
	if err != nil {
		logx.Error("[LoginLogic] getJwtToken error: ", err, " username: ", req.Username)
		return nil, errors.New(1002, "token error")
	}

	// redis 有token，直接返回
	_, err = l.svcCtx.RdsServer.Get(l.ctx, &db.GetRequest{Key: req.Username})
	if err == nil {
		return &types.LoginRsp{
			Token:   token,
			RetCode: 0,
		}, nil
	}

	// 交给rpc/chat 服务处理
	_, err = l.svcCtx.ChatClient.Login(l.ctx, &chat.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Error("[LoginLogic] rpc.Login error: ", err)
		return nil, errors.New(1001, "login failed")
	}

	logx.Infof("[LoginLogic] login success. username:%s token: %s", req.Username, token)
	return &types.LoginRsp{
		Token:   token,
		RetCode: 0,
	}, nil

}
