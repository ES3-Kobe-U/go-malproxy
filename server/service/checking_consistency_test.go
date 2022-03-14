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

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if err := CheckingTheIntegrityOfAmazonInformation(testcase.email, testcase.password); err != nil {
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

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if err := CheckingTheIntegrityOfRakutenInformation(testcase.userId, testcase.password); err != nil {
				log.Fatal(err)
			}
		})
	}
}
