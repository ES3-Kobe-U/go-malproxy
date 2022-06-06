package usecase

import (
	"context"
	"log"
	"testing"
)

func TestCheckingTheIntegrityOfAmazonInformation(t *testing.T) {
	amazonConfig, err := LoadAmazonEnv()
	if err != nil {
		t.Errorf("err:%v", err)
	}
	var usecases Usecase = &Contents{false, false}
	if usecases == nil {
		t.Errorf("services -> nil")
	}
	testcases := []struct {
		name     string
		ctx      context.Context
		email    string
		password string
	}{
		//各々テストケースを記述
		{
			name:     "example-00",
			ctx:      context.Background(),
			email:    amazonConfig.Email,
			password: amazonConfig.Password,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			_, err := usecases.CheckingTheIntegrityOfAmazonInformation(testcase.ctx, testcase.email, testcase.password)
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestCheckingTheIntegrityOfRakutenInformation(t *testing.T) {
	rakutenConfig, err := LoadRakutenEnv()
	var usecases Usecase
	if err != nil {
		t.Errorf("err:%v", err)
	}
	testcases := []struct {
		name     string
		ctx      context.Context
		userId   string
		password string
	}{
		//各々テストケースを記述
		{
			name:     "example-00",
			ctx:      context.Background(),
			userId:   rakutenConfig.UserId,
			password: rakutenConfig.Password,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			_, err := usecases.CheckingTheIntegrityOfRakutenInformation(testcase.ctx, testcase.userId, testcase.password)
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}
