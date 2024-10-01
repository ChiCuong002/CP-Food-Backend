package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	userPublicRouter := Router.Group("/user")
	{
		userPublicRouter.POST("/register")
		userPublicRouter.POST("/login")
		userPublicRouter.POST("/logout")
	}
	//private router
	userPrivateRouter := Router.Group("/user")
	{
		userPrivateRouter.GET("profile")
	}
}