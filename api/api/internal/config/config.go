/*
 * @Author: Leeson
 * @Date: 2024-11-29 17:51:51
 */
package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
