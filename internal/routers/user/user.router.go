package user

import (
	"food-recipes-backend/internal/middlewares"
	"food-recipes-backend/internal/wire"
	apierror "food-recipes-backend/pkg/errors"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//user controller
	userController, _ := wire.InitUserRouterHandler()
	//public router
	userPublicRouter := Router.Group("/user")
	{
		userPublicRouter.POST("/register", apierror.Make(userController.Register))
		userPublicRouter.POST("/login", apierror.Make(userController.Login))
	}
	// private router
	userPrivateRouter := Router.Group("/user")
	{
		userPrivateRouter.Use(middlewares.AuthMiddleware())
		userPrivateRouter.POST("/logout", apierror.Make(userController.Logout))
		userPrivateRouter.POST("/refresh-token", apierror.Make(userController.RefreshToken))
		userPrivateRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
