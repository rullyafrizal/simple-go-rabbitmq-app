package utils

import "os"

func LookupEnv(key string, fallback string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	return fallback
}
