package common

import "syscall"

func EnvString(key, fallback string) string {
	if envValue, ok := syscall.Getenv(key); ok {
		return envValue
	}
	// fallback can call itself a default value
	return fallback
}
