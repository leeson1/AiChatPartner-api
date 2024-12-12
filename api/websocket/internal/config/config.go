/*
 * @Author: LEESON
 * @Date: 2024-11-29 17:52:01
 */
package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
