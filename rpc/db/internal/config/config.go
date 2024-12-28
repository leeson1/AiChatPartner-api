/*
 * @Author: Leeson
 * @Date: 2024-12-15 12:31:10
 */
package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource  string
	RedisConfig redis.RedisConf
}
