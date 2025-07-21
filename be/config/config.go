package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

type Config struct {
	DBUri	string
	JWTSecret string
	AllowedOrigins []string
	Environment string
}

func Load() Config {
	// load .env only for local/dev environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	var allowedOrigins []string
	// prod domains
	if env == "production" {
		allowedOrigins = []string{
			"https://blackout.pages.dev",
		}
	} else {
		// dev domains
		allowedOrigins = []string{
			"http://localhost:5173",
		}
	}

	return Config {
		DBUri:	os.Getenv("MONGODB_URI"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		AllowedOrigins: allowedOrigins,
		Environment: env,
	}
}