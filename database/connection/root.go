package connection

import (
	"GoMVCStarterKit/config"
	"GoMVCStarterKit/database/connection/postgres"
	"strings"

	"xorm.io/xorm"
)

type SQLConnection interface {
	Connection(cfg *config.EnvConfig) *xorm.Session
}

func NewSQLDBConn(driver string) SQLConnection {
	switch strings.ToLower(driver) {
	case "mysql":
		return nil
	case "postgres":
		return &postgres.PostgreConn{}
	default:
		return nil
	}
}
