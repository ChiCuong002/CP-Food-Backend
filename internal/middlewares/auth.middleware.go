package middlewares

import (
	"food-recipes-backend/global"
	"food-recipes-backend/pkg/auth"
	"food-recipes-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        clientToken := c.Request.Header.Get("Authorization")
        if clientToken == "" {
            response.ErrorResponse(c, response.ErrorNoAuthorizationHeader, "No Authorization Header Provided")
            c.Abort()
            return
        }
        accessToken := clientToken[7:]
        claims := &auth.Claims{}
        token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(global.Config.Server.SecretKey), nil
        })
        if err != nil {
            if ve, ok := err.(*jwt.ValidationError); ok {
                if ve.Errors&jwt.ValidationErrorExpired != 0 {
                    response.ErrorResponse(c, response.ErrorUnauthorized, "Token has expired")
                } else {
                    response.ErrorResponse(c, response.ErrorUnauthorized, "Unauthorized")
                }
            } else {
                response.ErrorResponse(c, response.ErrorUnauthorized, "Unauthorized")
            }
            c.Abort()
            return
        }
        if !token.Valid {
            response.ErrorResponse(c, response.ErrorUnauthorized, "Unauthorized")
            c.Abort()
            return
        }
        c.Set("userId", claims.UserId)
        c.Next()
    }
}