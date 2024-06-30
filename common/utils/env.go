package utils

import (
	"syscall"

	"github.com/joho/godotenv"
)

func GetEnv(envFile, key, fallback string) string {
	err := godotenv.Load(envFile)
	if err != nil {
		return fallback
	}

	if value, ok := syscall.Getenv(key); ok {
		return value
	}

	return fallback
}

// don't know what perpose or when to use this function.
// just made it have more than GetEnv function
func SetEnv(key, value string) error {
	return syscall.Setenv(key, value)
}
