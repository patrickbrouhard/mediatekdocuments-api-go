package config

import "os"

type Config struct {
	AppEnv   string
	HTTPAddr string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() Config {
	return Config{
		AppEnv:   getEnv("APP_ENV", "development"),
		HTTPAddr: getEnv("HTTP_ADDR", ":8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "mediatek"),
		DBPassword: getEnv("DB_PASSWORD", "mediatek"),
		DBName:     getEnv("DB_NAME", "mediatek86"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
