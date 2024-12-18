/*
 * @Author: Leeson
 * @Date: 2024-12-09 17:38:22
 */
// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"AiChatPartner/api/websocket/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

type Server struct {
	svc *svc.ServiceContext
}

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	s := &Server{
		svc: serverCtx,
	}
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: s.WebsocketHandler(serverCtx),
			},
		},
	)
}
