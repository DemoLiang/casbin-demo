package svc

import (
	"github.com/DemoLiang/casbin-demo/internal/config"
	"github.com/DemoLiang/casbin-demo/internal/pkg/orm"
	"github.com/casbin/casbin/v2"
)

type ServiceContext struct {
	Config config.Config
	Cbn    *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	ormDB := orm.NewMysql(&orm.Config{
		DSN:    c.Mysql.DataSource,
		Active: c.Mysql.Active,
	})
	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher("mysql", ormDB, c.RedisConf)
	return &ServiceContext{
		Config: c,
		Cbn:    cbn,
	}
}
