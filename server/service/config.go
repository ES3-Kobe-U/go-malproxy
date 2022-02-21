package service

import (
	"os"

	"github.com/joho/godotenv"
)

type Rakuten struct {
	UserId   string
	Password string
}

type Amazon struct {
	Email    string
	Password string
}

func LoadRakutenEnv() (Rakuten, error) {
	var rakutenConfig = Rakuten{}
	err := godotenv.Load("data.env")
	if err != nil {
		return rakutenConfig, err
	}

	rakutenConfig.UserId = os.Getenv("RAKUTEN_USER_ID")
	rakutenConfig.Password = os.Getenv("RAKUTEN_PASSWORD")

	return rakutenConfig, nil
}

func LoadAmazonEnv() (Amazon, error) {
	var amazonConfig = Amazon{}
	err := godotenv.Load("data.env")
	if err != nil {
		return amazonConfig, err
	}

	amazonConfig.Email = os.Getenv("AMAZON_EMAIL")
	amazonConfig.Password = os.Getenv("AMAZON_PASSWORD")

	return amazonConfig, nil
}
