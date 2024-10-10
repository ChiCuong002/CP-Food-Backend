package middlewares

import (
	"errors"
	"fmt"
	"food-recipes-backend/global"
	"food-recipes-backend/pkg/auth"
	apierror "food-recipes-backend/pkg/errors"
	"strconv"
	"strings"

	//"food-recipes-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	AUTHORIZATION = "Authorization"
	CLIENT_ID     = "X-Client-Id"
	REFRESHTOKEN  = "X-Refresh-Token"
)
func verifyAccessToken(tokenString string) (*auth.Claims, error) {
    claims := &auth.Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(global.Config.Server.SecretKey), nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
                return nil, errors.New("token has expired")
            }
            return nil, errors.New("unauthorized")
        }
        return nil, errors.New("unauthorized")
    }
    if !token.Valid {
        return nil, errors.New("invalid access token")
    }
    return claims, nil
}
func verifyRefreshToken(tokenString string) (*auth.Claims, error) {
    claims := &auth.Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(global.Config.Server.SecretKey), nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
			fmt.Println("Validation Error Type: ", ve.Errors)
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
                return nil, errors.New("refresh token has expired")
            }
            return nil, errors.New("unauthorized")
        }
        return nil, errors.New("unauthorized")
    }
    if !token.Valid {
        return nil, errors.New("invalid refresh token")
    }
    return claims, nil
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := c.Request.Header.Get("X-Client-Id")
        if clientId == "" {
            apierror.ErrorHandler(c, apierror.NewAPIError(401, "Client ID is required"))
            c.Abort()
            return
        }

        authHeader := c.Request.Header.Get("Authorization")
        if authHeader == "" {
            apierror.ErrorHandler(c, apierror.NewAPIError(401, "Authorization token is required"))
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            apierror.ErrorHandler(c, apierror.NewAPIError(401, "Invalid token format"))
            c.Abort()
            return
        }

        var claims *auth.Claims
        var err error

        // Determine if the token is an access token or a refresh token
        if strings.Contains(c.Request.URL.Path, "/refresh-token") {
            claims, err = verifyRefreshToken(tokenString)
			c.Set("refreshToken", tokenString)
        } else {
            claims, err = verifyAccessToken(tokenString)
        }

        if err != nil {
            apierror.ErrorHandler(c, apierror.NewAPIError(401, err.Error()))
            c.Abort()
            return
        }

        clientIdInt, err := strconv.Atoi(clientId)
        if err != nil {
            apierror.ErrorHandler(c, apierror.NewAPIError(401, "Invalid Client ID"))
            c.Abort()
            return
        }

        if claims.UserId != clientIdInt {
            apierror.ErrorHandler(c, apierror.NewAPIError(401, "Invalid request"))
            c.Abort()
            return
        }

        fmt.Sprintln("Claims User ID: ", claims.UserId)
        c.Set("userId", claims.UserId)
        c.Next()
	}

}