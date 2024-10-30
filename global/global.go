package global

import (
	"database/sql"
	"food-recipes-backend/pkg/logger"
	"food-recipes-backend/pkg/setting"

	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.ZapLogger
	Db    *sql.DB
	Redis *redis.Client
)

func ProvideDB() *sql.DB {
    return Db
}

func ProvideRedis() *redis.Client {
	return Redis
}