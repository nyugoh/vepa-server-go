package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GoDotEnvVariable is...
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")

	}
	return os.Getenv(key)
}
