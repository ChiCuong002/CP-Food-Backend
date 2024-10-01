package initialize

import (
	//"food-recipes-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	r := gin.Default()
	//userController, _ := wire.InitUserRouterHandler()
	// use middleware
	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

	}
	return r
}