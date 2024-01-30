package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Locale         string
	FallbackLocale string
	PathLocale     string
}

// LoadConfig loads configuration from .env file and populates the Config struct.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	return &Config{
		Locale:         getEnv("TRANSLATION_LOCALE", "en"),
		FallbackLocale: getEnv("TRANSLATION_FALLBACK_LOCALE", "en"),
		PathLocale:     getEnv("TRANSLATION_PATH_LOCALE", "translation"),
	}, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		err := os.Setenv(key, fallback)
		if err != nil {
			return ""
		}
		return fallback
	}

	return value
}
