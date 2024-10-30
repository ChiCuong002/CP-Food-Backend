package admin

import (
	"food-recipes-backend/internal/wire"
	apierror "food-recipes-backend/pkg/errors"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminController, _ := wire.InitAdminRouterHandler()
	adminRouter := Router.Group("/admin")
	{
		adminRouter.POST("/login")
		adminRouter.POST("/logout")
		adminRouter.GET("/users", apierror.Make(adminController.ListUsers))
		adminRouter.GET("/user/:id", apierror.Make(adminController.DetailUser))
	}
}
