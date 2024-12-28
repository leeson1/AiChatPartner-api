/*
 * @Author: Leeson
 * @Date: 2024-11-24 00:37:56
 */
package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Dbserver zrpc.RpcClientConf
	AuthConf struct {
		AccessSecret string
		AccessExpire int64
	}
}
