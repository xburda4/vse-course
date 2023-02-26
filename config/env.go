package config

import (
	"os"
)

var (
	Port = "8080"
)

func init() {
	Port = GetEnv("PORT", Port)
}

// Config loads the environment config variable, or returns defaultValue if it
// is not set.
func GetEnv(name, defaultValue string) string {
	v := os.Getenv(name)
	if v != "" {
		return v
	}
	return defaultValue
}
