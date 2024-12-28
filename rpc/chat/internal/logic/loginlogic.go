/*
 * @Author: Leeson
 * @Date: 2024-12-14 17:25:47
 */
package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/chat/internal/svc"
	"AiChatPartner/rpc/db/db"

	"github.com/golang-jwt/jwt/v4"
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

func getJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(in *chat.LoginReq) (*chat.LoginRsp, error) {

	// 查db
	dbrsp, err := l.svcCtx.DbServer.Read(l.ctx, &db.ReadRequest{
		TableName: "ac_user",
		Key:       in.Username,
		KeyType:   2,
	})
	if err != nil {
		return nil, fmt.Errorf("[rpc/chat Login] get user by username:[%s] error: %s", in.Username, err)
	}

	pass := dbrsp.Data["password"]
	if in.Password != pass {
		return nil, fmt.Errorf("[rpc/chat Login] user:[%s] password error. ", in.Username)
	}

	//生成token
	now := jwt.TimeFunc().Unix()
	token, err := getJwtToken(l.svcCtx.Config.AuthConf.AccessSecret, now, l.svcCtx.Config.AuthConf.AccessExpire, in.Username)
	if err != nil {
		logx.Error("[LoginLogic] getJwtToken error: ", err, " username: ", in.Username)
		return nil, fmt.Errorf("[rpc/chat Login] user:[%s] get token error. ", in.Username)
	}
	dbrsp.Data["token"] = token

	// 插入redis
	jsonData, err := json.Marshal(dbrsp.Data)
	if err != nil {
		return &chat.LoginRsp{RetCode: 1}, fmt.Errorf("[rpc/chat Login] Error marshalling map: %s", err)
	}
	_, err = l.svcCtx.RdsServer.Set(l.ctx, &db.SetRequest{
		Key:   in.Username,
		Value: string(jsonData),
	})
	if err != nil {
		return &chat.LoginRsp{RetCode: 1}, fmt.Errorf("[rpc/chat Login] Error setting key: %s", err)
	}

	logx.Info("[rpc/chat Login] login success. user: ", in.Username)

	return &chat.LoginRsp{RetCode: 0, Data: token}, nil
}
