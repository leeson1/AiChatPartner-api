/*
 * @Author: Leeson
 * @Date: 2024-12-15 15:01:46
 */
package redisservicelogic

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *db.GetRequest) (*db.GetResponse, error) {

	value, err := l.svcCtx.RedisClient.GetCtx(l.ctx, in.Key)
	if err != nil {
		return &db.GetResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &db.GetResponse{
		Success: true,
		Value:   value,
	}, nil
}
