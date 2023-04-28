package service

import (
	"context"
	"customer/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtCustomClaim struct {
	ID int `json:"id"`
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

func JwtGenerate(ctx context.Context, userID int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
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