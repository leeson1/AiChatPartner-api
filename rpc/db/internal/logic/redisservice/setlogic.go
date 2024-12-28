/*
 * @Author: Leeson
 * @Date: 2024-12-29 01:28:51
 */
package redisservicelogic

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLogic {
	return &SetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetLogic) Set(in *db.SetRequest) (*db.SetResponse, error) {

	err := l.svcCtx.RedisClient.SetCtx(l.ctx, in.Key, in.Value)
	if err != nil {
		return &db.SetResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}

	return &db.SetResponse{Success: true}, nil
}
