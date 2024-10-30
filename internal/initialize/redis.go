package initialize

import (
	"fmt"
	"food-recipes-backend/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	config := global.Config.Redis
	url := fmt.Sprintf("redis://%s:%s@%s/%d", config.User, config.Password, config.Addr, config.DB)
	opt, err := redis.ParseURL(url)
	if err != nil {
		fmt.Println("failed to parse redis url", err)
		global.Logger.Error("failed to parse redis url", zap.Error(err))
	}
	rdb := redis.NewClient(opt)
	global.Redis = rdb
	global.Logger.Info("Successfully connected to redis")
}