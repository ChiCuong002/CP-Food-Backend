package middlewares

import (
	"fmt"
	"food-recipes-backend/global"
	apierror "food-recipes-backend/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		redisClient := global.ProvideRedis()
		serveLimit := global.Config.Server.RateLimit
		limitDuration := global.Config.Server.RateLimitDuration
		identifier := c.ClientIP()
		fmt.Println("server limit: ", serveLimit)	
		key := fmt.Sprintf("rate_limit:%s", identifier)
		
		// increase the counter
		count, err := redisClient.Incr(c, key).Result()
		if err != nil {
			apierror.ErrorHandler(c, apierror.NewAPIError(500, "Redis server error"))
            c.Abort()
            return
		}
		if count == 1 {
			redisClient.Expire(c, key, time.Duration(limitDuration) * time.Second)
		}
		if count > int64(serveLimit) {
			apierror.ErrorHandler(c, apierror.NewAPIError(429, "Rate limit exceeded"))
			c.Abort()
			return
		}

		c.Next()
	}
}