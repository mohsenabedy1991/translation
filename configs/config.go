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
		Locale:         os.Getenv("TRANSLATION_LOCALE"),
		FallbackLocale: os.Getenv("TRANSLATION_FALLBACK_LOCALE"),
		PathLocale:     os.Getenv("TRANSLATION_PATH_LOCALE"),
	}, nil
}
