/*
 * @Author: Leeson
 * @Date: 2024-12-01 17:06:26
 */
package logic

import (
	"context"

	"AiChatPartner/api/api/internal/svc"
	"AiChatPartner/api/api/internal/types"

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
	// todo: add your logic here and delete this line

	if req.Username == "admin" && req.Password == "admin" {
		//生成token
		now := jwt.TimeFunc().Unix()
		token, err := getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, req.Username)
		if err != nil {
			return nil, err
		}
		logx.Info("[LoginLogic] login success. token: ", token)
		return &types.LoginRsp{
			Token:   token,
			RetCode: 200,
		}, nil
	}

	return
}
