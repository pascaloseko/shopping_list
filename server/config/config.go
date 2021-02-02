package config

import (
	"os"
)

// Config func to get env value from key ---
func Config(key string) string {
	return os.Getenv(key)
}
