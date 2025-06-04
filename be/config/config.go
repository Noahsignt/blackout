package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

type Config struct {
	DBUri	string
}

func Load() Config {
	// load .env only for local/dev environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	return Config {
		DBUri:	os.Getenv("MONGODB_URI"),
	}
}