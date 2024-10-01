package routers

import (
	"food-recipes-backend/global"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	// r.Use() // logger
	// r.Use() // cors	
	// r.Use() // limit request

	// routers
	userRouter := RouterGroupApp.User
	MainGroup := r.Group("/v1/api")
	{
		userRouter.InitUserRouter(MainGroup)
	}
	return r
}