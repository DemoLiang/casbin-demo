package orm

import (
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN         string
	Active      int
	Idle        int
	IdleTimeout time.Duration
	LogxConf    logx.LogConf
}

func NewMysql(c *Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{
		Logger: NewGormLogger(c.LogxConf),
	})
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		logx.Errorf("db set idle conn err :%v", err)
	}

	sqlDB.SetMaxIdleConns(c.Idle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(c.Active)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(c.IdleTimeout)
	// TODO 实现缓存
	//db.Plugins["cache"]
	//db.Use("cachePlugin")

	db.Exec("SET NAMES utf8mb4;")

	return db
}
