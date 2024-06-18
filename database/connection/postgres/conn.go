package postgres

import (
	"GoMVCStarterKit/config"
	"GoMVCStarterKit/src/utils/logger"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type PostgreConn struct{}

func (p *PostgreConn) Connection(cfg *config.EnvConfig) *xorm.Session {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Hostname,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("failed create engine xorm got : %v", err))
	}
	engine.SetMaxIdleConns(cfg.Database.MaxIdleConn)
	engine.SetMaxOpenConns(cfg.Database.MaxOpenConn)
	session := engine.NewSession()
	return session
}
