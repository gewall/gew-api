package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	err := godotenv.Load(".env." + env + ".local")

	if err != nil {
		return err
	}

	fmt.Println("Server run on env", env)

	return nil
}
