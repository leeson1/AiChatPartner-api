/*
 * @Author: LEESON
 * @Date: 2024-11-29 17:51:51
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"AiChatPartner/api/api/internal/config"
	"AiChatPartner/api/api/internal/handler"
	"AiChatPartner/api/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/api-api.yaml", "the config file")
	c          config.Config
)

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	logx.AddWriter(logx.NewWriter(os.Stdout))

	server := rest.MustNewServer(c.RestConf, rest.WithCorsHeaders("loginHandler", "userInfoHandler", "registerHandler"))
	// server.Use(middle.Middleware)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
