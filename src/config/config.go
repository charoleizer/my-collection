package config

import (
	"github.com/joho/godotenv"
)

func LoadDatabaseEnv() error {
	err := godotenv.Load("src/config/config.env")
	if err != nil {
		return err
	}

	return nil
}
