package routers

import (
	"food-recipes-backend/global"
	"food-recipes-backend/internal/middlewares"
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:     true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	r.Use(middlewares.RateLimitMiddleware())

	// routers
	userRouter := RouterGroupApp.User
	adminRouter := RouterGroupApp.Admin
	MainGroup := r.Group("/v1/api")
	{
		userRouter.InitUserRouter(MainGroup)
		adminRouter.InitAdminRouter(MainGroup)
	}
	return r
}