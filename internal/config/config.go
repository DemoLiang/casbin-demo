package config

import (
	"github.com/DemoLiang/casbin-demo/internal/pkg/casbin"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RedisConf  redis.RedisConf
	CasbinConf casbin.CasbinConf
	Mysql      struct {
		DataSource  string
		Active      int
		Idle        int
		IdleTimeout int
	}
}
