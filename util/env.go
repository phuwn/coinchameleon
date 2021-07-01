package util

import "os"

func Getenv(key string, fallback string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return fallback
}
