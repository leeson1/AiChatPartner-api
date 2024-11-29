/*
 * @Author: LEESON
 * @Date: 2024-11-24 14:16:10
 */

package handler

import (
	"net/http"

	"AiChatPartner/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/expand",
				Handler: ExpandHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/shorten",
				Handler: ShortenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: WebsocketHandler(serverCtx),
			},
		},
	)
}
