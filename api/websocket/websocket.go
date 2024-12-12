/*
 * @Author: LEESON
 * @Date: 2024-11-29 17:52:01
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"AiChatPartner/api/websocket/internal/config"
	"AiChatPartner/api/websocket/internal/handler"
	"AiChatPartner/api/websocket/internal/svc"
	"AiChatPartner/common"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/websocket.yaml", "the config file")
	c          config.Config
)

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	logx.AddWriter(logx.NewWriter(os.Stdout))

	if err := common.InitServices("../../common/etc/common.yaml"); err != nil {
		logx.Field("failed to initialize services: %v", err)
	}

	s := rest.MustNewServer(c.RestConf)
	defer s.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(s, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	s.Start()
}
