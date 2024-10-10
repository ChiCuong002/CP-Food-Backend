package initialize

import (
	"food-recipes-backend/internal/routers"
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitDatabase()
	r := routers.InitializeRoutes()
	return r
}
