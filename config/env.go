package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBPass               string
	DBAddr               string
	DBName               string
	JwtExpirationSeconds int64
	JwtSecret            string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBUser:               getEnv("DB_USER", "ecom"),
		DBPass:               getEnv("DB_PASS", "password"),
		DBAddr:               fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:               getEnv("DB_NAME", "ecom"),
		JwtExpirationSeconds: getEnvAsInt("JWT_EXPIRATION_SECONDS", 3600*24*7), // 7 days
		JwtSecret:            getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return fallback
		}
		return intValue
	}
	return fallback
}
