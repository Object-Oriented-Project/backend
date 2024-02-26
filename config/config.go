package config

import (
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}