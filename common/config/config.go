/*
 * @Author: LEESON
 * @Date: 2024-12-12 15:38:54
 */
// common/config/config.go
package config

import (
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQLConf MySQLConfig `yaml:"MySQLConf"`
	RedisConf RedisConfig `yaml:"RedisConf"`
}

type MySQLConfig struct {
	// MySQL 配置字段
	Host     string `yaml:"Host"`
	Username string `yaml:"Username"`
	Pass     string `yaml:"Pass"`
}

type RedisConfig struct {
	// Redis 配置字段
	Host string `yaml:"Host"`
	Pass string `yaml:"Pass"`
	Type string `yaml:"Type"`
	Tls  bool   `yaml:"Tls"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var c Config
	// err = yaml.Unmarshal(data, &c)
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		logx.Error("[LoadConfig] err:", err)
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	logx.Info("[LoadConfig] common data:", string(data))

	return &c, nil
}
