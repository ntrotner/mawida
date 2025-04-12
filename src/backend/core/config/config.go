package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var GlobalConfig *Config

type AppEnvironment string

const (
	Development AppEnvironment = "dev"
	Production  AppEnvironment = "prod"
)

// Config holds all configuration for the application
type Config struct {
	Environment AppEnvironment
	Server      ServerConfig
	Database    DatabaseConfig
	Auth        AuthConfig
}

// ServerConfig holds all server related configuration
type ServerConfig struct {
	Address string
}

// DatabaseConfig holds all database related configuration
type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
}

// AuthConfig holds all authentication related configuration
type AuthConfig struct {
	JWTSecret              string
	TokenExpirationMinutes int
}

// Load loads all configuration from environment variables
func Load() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found, using environment variables")
	}

	// Set default token expiration to 30 minutes if not specified
	tokenExpiration := 30
	if expStr := os.Getenv("jwtExpirationMinutes"); expStr != "" {
		exp, err := strconv.Atoi(expStr)
		if err == nil {
			tokenExpiration = exp
		}
	}

	GlobalConfig = &Config{
		Environment: determineEnvironment(),
		Server: ServerConfig{
			Address: getEnv("host", "0.0.0.0:8080"),
		},
		Database: DatabaseConfig{
			Host: getEnv("databaseHost", "http://admin:password@localhost:5984"),
		},
		Auth: AuthConfig{
			JWTSecret:              getEnv("jwtKey", "test-secret"),
			TokenExpirationMinutes: tokenExpiration,
		},
	}
}

// determineEnvironment determines the environment of the application
func determineEnvironment() AppEnvironment {
	environment, exists := os.LookupEnv("environment")
	if !exists {
		return Production
	}

	switch environment {
	case string(Development):
		return Development
	default:
		return Production
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
