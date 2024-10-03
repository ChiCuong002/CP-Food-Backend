package user

import (
	"food-recipes-backend/internal/middlewares"
	"food-recipes-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//user controller
	userController, _ := wire.InitUserRouterHandler()
	//public router
	userPublicRouter := Router.Group("/user")
	{
		userPublicRouter.POST("/register", userController.Register)
		userPublicRouter.POST("/login", userController.Login)
		userPublicRouter.POST("/logout")
	}
	//private router
	userPrivateRouter := Router.Group("/user")
	{
		userPrivateRouter.Use(middlewares.TokenAuthMiddleware())	
		userPrivateRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}