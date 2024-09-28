package config

import (
	"log"
	"os"
)

var (
	AppName  string
	LogLevel string
)

func LoadConfig() {
	AppName = getEnv("APP_NAME", "Startup")
	LogLevel = getEnv("LOG_LEVEL", "debug")

	log.Printf("Configuration loaded: AppName=%s, LogLevel=%s", AppName, LogLevel)
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
