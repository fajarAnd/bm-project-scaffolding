package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
	Storage  StorageConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
}

type JWTConfig struct {
	Secret string
	Expiry string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type StorageConfig struct {
	Type   string // "local" or "s3"
	Bucket string
	Region string
}

func Load() (*Config, error) {
	// Load .env file in development
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load()
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			URL:          getEnv("DATABASE_URL", ""),
			MaxOpenConns: 25,
			MaxIdleConns: 5,
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""),
			Expiry: getEnv("JWT_EXPIRY", "24h"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},
		Storage: StorageConfig{
			Type:   getEnv("STORAGE_TYPE", "local"),
			Bucket: getEnv("S3_BUCKET", ""),
			Region: getEnv("S3_REGION", "us-east-1"),
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.Database.URL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}