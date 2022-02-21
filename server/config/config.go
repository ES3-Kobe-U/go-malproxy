package config

import (
	"os"

	"github.com/joho/godotenv"
)

type rakuten struct {
	userId   string
	password string
}

var rakutenConfig = rakuten{}

type amazon struct {
	email    string
	password string
}

var amazonConfig = amazon{}

func loadEnv() error {
	err := godotenv.Load("data.env")
	if err != nil {
		return err
	}
	{
		rakutenConfig.userId = os.Getenv("RAKUTEN_USER_ID")
		rakutenConfig.password = os.Getenv("RAKUTEN_PASSWORD")
		amazonConfig.email = os.Getenv("AMAZON_EMAIL")
		amazonConfig.password = os.Getenv("AMAZON_PASSWORD")
	}
	return nil
}
