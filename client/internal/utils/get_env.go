package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("init env has failed failed with error: %v\n", err)
		os.Exit(1)
	}
}
