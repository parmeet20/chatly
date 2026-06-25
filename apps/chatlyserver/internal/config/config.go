package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	PORT                      string
	AllowedOrigins            []string
	MONGO_URL                 string
	JWT_SECRET                string
	JWT_TOKEN_EXPIRATION_TIME time.Duration
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT is required")
	}

	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		return nil, fmt.Errorf("MONGO_URL is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	allowedOrigins := parseOrigins(os.Getenv("ALLOWED_ORIGINS"))
	if len(allowedOrigins) == 0 {
		return nil, fmt.Errorf("ALLOWED_ORIGINS is required")
	}

	jwtExpirationStr := os.Getenv("JWT_TOKEN_EXPIRATION_TIME")
	if jwtExpirationStr == "" {
		return nil, fmt.Errorf("JWT_TOKEN_EXPIRATION_TIME is required")
	}

	jwtExpiration, err := time.ParseDuration(jwtExpirationStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_TOKEN_EXPIRATION_TIME: %w", err)
	}

	return &Config{
		PORT:                      port,
		MONGO_URL:                 mongoURL,
		JWT_SECRET:                jwtSecret,
		AllowedOrigins:            allowedOrigins,
		JWT_TOKEN_EXPIRATION_TIME: jwtExpiration,
	}, nil
}

func parseOrigins(origins string) []string {

	parts := strings.Split(origins, ",")

	var cleaned []string
	for _, o := range parts {
		o = strings.TrimSpace(o)
		if o != "" {
			cleaned = append(cleaned, o)
		}
	}

	return cleaned
}
