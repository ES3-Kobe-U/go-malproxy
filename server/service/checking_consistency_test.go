package service

import (
	"log"
	"testing"
)

func TestCheckingTheIntegrityOfAmazonInformation(t *testing.T) {
	amazonConfig, err := LoadAmazonEnv()
	if err != nil {
		t.Errorf("err:%v", err)
	}
	testcases := []struct {
		name     string
		email    string
		password string
	}{
		//各々テストケースを記述
		{
			name:     "example-00",
			email:    amazonConfig.Email,
			password: amazonConfig.Password,
		},
	}

	for _, tetestcase := range testcases {
		t.Run(tetestcase.name, func(t *testing.T) {
			if err := CheckingTheIntegrityOfAmazonInformation(tetestcase.email, tetestcase.password); err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestCheckingTheIntegrityOfRakutenInformation(t *testing.T) {
	rakutenConfig, err := LoadRakutenEnv()
	if err != nil {
		t.Errorf("err:%v", err)
	}
	testcases := []struct {
		name     string
		userId   string
		password string
	}{
		//各々テストケースを記述
		{
			name:     "example-00",
			userId:   rakutenConfig.UserId,
			password: rakutenConfig.Password,
		},
	}

	for _, tetestcase := range testcases {
		t.Run(tetestcase.name, func(t *testing.T) {
			if err := CheckingTheIntegrityOfRakutenInformation(tetestcase.userId, tetestcase.password); err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestGoRakuten(t *testing.T) {
	if err := GoRakuten(); err != nil {
		t.Errorf("err:%v", err)
	}
}
