package config

import (
	"os"

	"github.com/gin-contrib/cors"
)

type AppConfig struct {
	GetSongAPIKey string
	AllowedOrigin string
}

func Load() AppConfig {
	return AppConfig{
		GetSongAPIKey: os.Getenv("GETSONG_API_KEY"),
		AllowedOrigin: os.Getenv("CORS_ORIGIN"),
	}
}

func CORS(cfg AppConfig) cors.Config {
	return cors.Config{
		AllowOrigins:     []string{cfg.AllowedOrigin},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}
}
