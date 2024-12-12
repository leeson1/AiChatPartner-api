/*
 * @Author: LEESON
 * @Date: 2024-12-12 15:40:56
 */
package common

import (
	"AiChatPartner/common/config"
	"AiChatPartner/common/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

func InitServices(configPath string) error {
	logx.Info("[InitServices] configPath:", configPath)
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	// if err := mysql.InitMySQL(conf.MySQL); err != nil {
	// 	return err
	// }
	if err := redis.InitRedis(conf); err != nil {
		return err
	}

	return nil
}
