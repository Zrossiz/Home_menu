package config

import "os"

type Config struct {
	DB struct {
		DBURI string
	}
	Logger struct {
		Level string
	}
	Application struct {
		Port      string
		ClientURL string
	}
}

func GetConfig() Config {
	var cfg Config
	cfg.Application.Port = getStringEnvOrDefault("PORT", ":8080")
	cfg.Application.ClientURL = getStringEnvOrDefault("CLIENT_URL", "http://localhost:5173")
	cfg.DB.DBURI = getStringEnvOrDefault("DB_URI", "")
	cfg.Logger.Level = getStringEnvOrDefault("LOG_LEVEL", "debug")

	return cfg
}

func getStringEnvOrDefault(envName, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		return defaultValue
	}

	return envValue
}
