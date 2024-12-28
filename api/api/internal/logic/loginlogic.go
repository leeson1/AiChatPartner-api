/*
 * @Author: Leeson
 * @Date: 2024-12-01 17:06:26
 */
package logic

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/x/errors"

	"AiChatPartner/api/api/internal/svc"
	"AiChatPartner/api/api/internal/types"
	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/db/db"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {

	// redis 有token，直接返回
	rdsRsp, err := l.svcCtx.RdsServer.Get(l.ctx, &db.GetRequest{Key: req.Username})
	if err == nil {
		jsonData := []byte(rdsRsp.Value)
		var data map[string]interface{}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			logx.Error("[LoginLogic] json.Unmarshal error: ", err)
			return nil, errors.New(1003, "json.Unmarshal error")
		}

		return &types.LoginRsp{
			Token:   data["token"].(string),
			RetCode: 0,
		}, nil
	}

	// 无token, 交给rpc/chat 服务处理
	chatRsp, err := l.svcCtx.ChatClient.Login(l.ctx, &chat.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Error("[LoginLogic] rpc.Login error: ", err)
		return nil, errors.New(1001, "login failed")
	}

	logx.Infof("[LoginLogic] login success. username:%s token: %s", req.Username, chatRsp.Data)
	return &types.LoginRsp{
		Token:   chatRsp.Data,
		RetCode: 0,
	}, nil

}
