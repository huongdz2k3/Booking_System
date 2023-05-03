package service

import (
	"context"
	"customer/internal/logger"
	"customer/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtCustomClaim struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	secret := get("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}

func JwtGenerate(userID int, role string) (string, error) {
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	timeExpired, err := time.ParseDuration(get("JWT_EXPIRE"))

	if err != nil {
		logger.NewLogger().Error("JWT_EXPIRE is not a number")
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID:   userID,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * timeExpired).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	token, err := t.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.WrapGQLBadRequestError(ctx, "there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
