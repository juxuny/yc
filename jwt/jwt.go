package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func SetSecret(secret string) {
	jwtSecret = []byte(secret)
}

var jwtSecret = []byte("")

type Claims struct {
	UserId   uint64 `json:"uid"`
	UserName string `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint64, userName string) (string, error) {
	if len(jwtSecret) == 0 {
		return "", fmt.Errorf("secret is empty")
	}
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "yc",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
