package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("/admin")
	{
		adminRouter.POST("login")
		adminRouter.POST("logout")
	}
}
	