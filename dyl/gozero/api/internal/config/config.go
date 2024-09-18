package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}
	Auth AuthConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
