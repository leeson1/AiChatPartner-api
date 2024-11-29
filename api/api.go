/*
 * @Author: LEESON
 * @Date: 2024-11-24 00:37:56
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"AiChatPartner/api/internal/config"
	"AiChatPartner/api/internal/handler"
	"AiChatPartner/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	logx.AddWriter(logx.NewWriter(os.Stdout))

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
