package svc

import (
	"github.com/DemoLiang/casbin-demo/internal/config"
	"github.com/casbin/casbin/v2"
)

type ServiceContext struct {
	Config config.Config
	Cbn    *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher("mysql", c.Mysql.DataSource, c.RedisConf, c.CasbinConf.ModelText)
	return &ServiceContext{
		Config: c,
		Cbn:    cbn,
	}
}
