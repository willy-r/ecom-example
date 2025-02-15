package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/willy-r/ecom-example/config"
)

func CreateJwt(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JwtExpirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
