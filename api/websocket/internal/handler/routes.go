/*
 * @Author: LEESON
 * @Date: 2024-11-29 17:52:01
 */

package handler

import (
	"net/http"

	"AiChatPartner/api/websocket/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: WebsocketHandler(serverCtx),
			},
		},
	)
}
