package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/willy-r/ecom-example/config"
	"github.com/willy-r/ecom-example/types"
	"github.com/willy-r/ecom-example/utils"
)

type ContextKey string

const UserKey ContextKey = "userId"

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

func WithJwtAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)

		token, err := ValidateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			utils.PermissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("invalid token")
			utils.PermissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)

		userId, _ := strconv.Atoi(str)

		u, err := store.GetUserById(userId)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			utils.PermissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth == "" || !strings.HasPrefix(tokenAuth, "Bearer ") {
		return ""
	}

	return strings.TrimPrefix(tokenAuth, "Bearer ")
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Envs.JwtSecret), nil
	})
}

func GetUserIdFromContext(ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userId
}
