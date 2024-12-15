/*
 * @Author: LEESON
 * @Date: 2024-12-12 15:40:56
 */
package common

import (
	"AiChatPartner/common/config"
	"AiChatPartner/common/redis"
)

func InitServices(configPath string) error {
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	// if err := mysql.InitMySQL(conf); err != nil {
	// 	return err
	// }
	if err := redis.InitRedis(conf); err != nil {
		return err
	}

	return nil
}
