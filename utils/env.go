package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	CheckError(err)
	value := os.Getenv(key)
	return value
}
