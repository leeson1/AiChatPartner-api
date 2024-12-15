/*
 * @Author: Leeson
 * @Date: 2024-12-15 14:38:28
 */
package logic

import (
	"context"

	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/chat/internal/svc"
	"AiChatPartner/rpc/db/db"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *chat.RegisterReq) (*chat.RegisterRsp, error) {

	// 1. 先拿username 查 redis，如果存在直接返回用户名重复 retCode = 1
	// userInfo, err := redis.GetRedisClient().Get("admin")
	// if err != nil {
	// 	logx.Error("[RegisterLogic] get user by username error: ", err)
	// 	return &chat.RegisterRsp{
	// 		RetCode: 2,
	// 	}, err
	// }
	// if userInfo != "" {
	// 	logx.Error("[RegisterLogic] user already exists. username: ", in.Username)
	// 	return &chat.RegisterRsp{
	// 		RetCode: 1,
	// 	}, fmt.Errorf("[RegisterLogic] user already exists. username: %s", in.Username)
	// }

	userInfo, err := l.svcCtx.DbServer.Read(l.ctx, &db.ReadRequest{
		TableName: "ac_user",
		Key:       in.Username,
		KeyType:   2,
	})

	var notFound bool = false
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if sqlx.ErrNotFound.Error() == st.Message() {
				notFound = true
			} else {
				return &chat.RegisterRsp{RetCode: 2}, err
			}
		} else {
			logx.Errorf("[rpc/chat Register] get user by username:[%s] error: [%s]", in.Username, err)
			return &chat.RegisterRsp{RetCode: 2}, err
		}
	}

	if userInfo.Success {
		logx.Errorf("[rpc/chat Register] username:[%s] already exists", in.Username)
		return &chat.RegisterRsp{RetCode: 2}, nil
	}

	// 2. 不存在，调用db服务写入数据
	if notFound {

		data := make(map[string]string)
		data["role"] = "2"
		data["username"] = in.Username
		data["password"] = in.Password

		_, err := l.svcCtx.DbServer.Create(l.ctx, &db.CreateRequest{
			TableName: "ac_user",
			Data:      data,
		})
		if err != nil {
			logx.Errorf("[rpc/chat Register] Create username:[%s] error: [%s]", in.Username, err)
			return &chat.RegisterRsp{RetCode: 1}, nil
		}
	}

	logx.Infof("[rpc/chat Register] register username:[%s] success. ", in.Username)

	return &chat.RegisterRsp{RetCode: 0}, nil
}
