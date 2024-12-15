/*
 * @Author: Leeson
 * @Date: 2024-12-15 12:31:10
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/config"
	"AiChatPartner/rpc/db/internal/server"
	"AiChatPartner/rpc/db/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/db.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	logx.AddWriter(logx.NewWriter(os.Stdout))

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		db.RegisterDatabaseServiceServer(grpcServer, server.NewDatabaseServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
