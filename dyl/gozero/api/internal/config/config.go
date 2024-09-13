package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth AuthConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
