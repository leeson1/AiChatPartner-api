// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: chat.proto

package chatclient

import (
	"context"

	"AiChatPartner/rpc/chat/chat"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ExpandReq   = chat.ExpandReq
	ExpandResp  = chat.ExpandResp
	Request     = chat.Request
	Response    = chat.Response
	ShortenReq  = chat.ShortenReq
	ShortenResp = chat.ShortenResp

	Chat interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error)
		Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := chat.NewChatClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultChat) Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error) {
	client := chat.NewChatClient(m.cli.Conn())
	return client.Expand(ctx, in, opts...)
}

func (m *defaultChat) Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error) {
	client := chat.NewChatClient(m.cli.Conn())
	return client.Shorten(ctx, in, opts...)
}
