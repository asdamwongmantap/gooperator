package config

import (
	"os"
)

func MyPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}
	return ":" + port, nil
}
