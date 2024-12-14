// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: chat.proto

package server

import (
	"context"

	"AiChatPartner/rpc/chat/chat"
	"AiChatPartner/rpc/chat/internal/logic"
	"AiChatPartner/rpc/chat/internal/svc"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) Login(ctx context.Context, in *chat.LoginReq) (*chat.LoginRsp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *ChatServer) UserInfo(ctx context.Context, in *chat.UserInfoReq) (*chat.UserInfoRsp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
