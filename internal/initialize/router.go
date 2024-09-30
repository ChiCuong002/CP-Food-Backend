package initialize

import "github.com/gin-gonic/gin"

func Initialize() *gin.Engine {
	r := gin.Default()
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