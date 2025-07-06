package auth

import (
	"time"

	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func GenerateTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken string) (string, error) {
	// 1. Set time expiration
	timeExpiration := global.Config.JWT.JWT_EXPIRATION
	if timeExpiration == "" {
		timeExpiration = "1h"
	}
	expiration, err := time.ParseDuration(timeExpiration)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expireAt := now.Add(expiration)
	return GenerateTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expireAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "shopdevgo",
			Subject:   uuidToken,
		},
	})
}
