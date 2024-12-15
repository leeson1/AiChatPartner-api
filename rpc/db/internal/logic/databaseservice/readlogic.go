/*
 * @Author: Leeson
 * @Date: 2024-12-15 15:01:46
 */
package databaseservicelogic

import (
	"context"
	"strconv"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/svc"
	"AiChatPartner/rpc/db/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadLogic {
	return &ReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadLogic) Read(in *db.ReadRequest) (*db.ReadResponse, error) {

	var userInfo *model.AcUser
	var err error

	if in.KeyType == 1 { //主键 id
		id, err := strconv.ParseUint(in.Key, 10, 64)
		if err != nil {
			return &db.ReadResponse{Success: false, ErrorMessage: err.Error()}, err
		}
		userInfo, err = l.svcCtx.Model.FindOne(l.ctx, id)
		if err != nil {
			return &db.ReadResponse{Success: false, ErrorMessage: err.Error()}, err
		}

	} else if in.KeyType == 2 { // 唯一索引 username
		userInfo, err = l.svcCtx.Model.GetUserByUsername(l.ctx, in.Key)
		if err != nil {
			return &db.ReadResponse{Success: false, ErrorMessage: err.Error()}, err
		}
	}

	userInfoMap := make(map[string]string)
	userInfoMap["id"] = strconv.FormatUint(userInfo.Uin, 10)
	userInfoMap["role"] = strconv.FormatUint(uint64(userInfo.Role), 10)
	userInfoMap["username"] = userInfo.Username
	userInfoMap["password"] = userInfo.Password
	if userInfo.Email.Valid {
		userInfoMap["email"] = userInfo.Email.String
	} else {
		userInfoMap["email"] = ""
	}
	if userInfo.Nickname.Valid {
		userInfoMap["nickname"] = userInfo.Nickname.String
	} else {
		userInfoMap["nickname"] = ""
	}
	if userInfo.Sex.Valid {
		userInfoMap["sex"] = userInfo.Sex.String
	} else {
		userInfoMap["sex"] = ""
	}
	if userInfo.CreateTime.Valid {
		userInfoMap["create_time"] = userInfo.CreateTime.Time.String()
	} else {
		userInfoMap["create_time"] = ""
	}
	if userInfo.UpdateTime.Valid {
		userInfoMap["update_time"] = userInfo.UpdateTime.Time.String()
	} else {
		userInfoMap["update_time"] = ""
	}
	if userInfo.Version.Valid {
		userInfoMap["version"] = strconv.FormatInt(userInfo.Version.Int64, 10)
	} else {
		userInfoMap["version"] = ""
	}

	return &db.ReadResponse{
		Success: true,
		Data:    userInfoMap,
	}, nil
}
