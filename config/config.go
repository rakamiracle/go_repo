package config

import (
    "os"
)

type Config struct {
    Port string
}

func GetConfig() *Config {
    return &Config{
        Port: getEnv("APP_PORT", "8080"),
    }
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
