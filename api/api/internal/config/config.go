/*
 * @Author: Leeson
 * @Date: 2024-11-29 17:51:51
 */
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Chat     zrpc.RpcClientConf
	Dbserver zrpc.RpcClientConf
	Auth     struct {
		AccessSecret string
		AccessExpire int64
	}
}
