package global

import (
	"database/sql"
	"food-recipes-backend/pkg/logger"
	"food-recipes-backend/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.ZapLogger
	Db    *sql.DB
)

func ProvideDB() *sql.DB {
    return Db
}