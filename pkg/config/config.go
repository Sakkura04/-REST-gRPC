package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file")
		os.Exit(1)
	}
}
