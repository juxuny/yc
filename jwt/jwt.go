package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/juxuny/yc/dt"
	"time"
)

func SetSecret(secret string) {
	jwtSecret = []byte(secret)
}

var jwtSecret = []byte("")

var tokenValidityPeriod = time.Hour * 24

func SetTokenValidityPeriod(duration time.Duration) {
	tokenValidityPeriod = duration
}

type Claims struct {
	UserId   dt.ID  `json:"uid"`
	UserName string `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(userId dt.ID, userName string) (string, error) {
	if len(jwtSecret) == 0 {
		return "", fmt.Errorf("secret is empty")
	}
	nowTime := time.Now()
	expireTime := nowTime.Add(tokenValidityPeriod)

	claims := Claims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprintf("%v", userId.Uint64),
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "YuanJie Cloud",
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
