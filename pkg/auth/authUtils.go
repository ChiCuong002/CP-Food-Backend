package auth

import (
	"fmt"
	"food-recipes-backend/global"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func CreateTokenPair(payload int) (string, string, error) {
	// set the expiration time of the token
	accessTokenExpirationTime := time.Now().Add(15 * time.Minute).Unix()
	refreshTokenExpirationTime := time.Now().Add(24 * time.Hour).Unix()

	// create JWT claims
	accessTokenClaims := &Claims{
		UserId: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpirationTime,
		},
	}
	refreshTokenClaims := &Claims{
		UserId: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpirationTime,
		},
	}

	// create token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	// sign the token with the secret key
	sercretKey := global.Config.Server.SecretKey
	accessTokenString, err := accessToken.SignedString([]byte(sercretKey))
	if err != nil {
		fmt.Println("Error signing access token")
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(sercretKey))
	if err != nil {
		fmt.Println("Error signing refresh token")
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}