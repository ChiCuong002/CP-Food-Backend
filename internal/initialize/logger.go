package initialize

import (
	"food-recipes-backend/global"
	"food-recipes-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}