package service

import (
	"log"
	"testing"
)

func TestCheckingTheIntegrityOfAmazonInformation(t *testing.T) {
	testcases := []struct {
		name     string
		email    string
		password string
	}{
		//各々テストケースを記述
		{
			name:     "example-00",
			email:    `example.00@email.com`,
			password: `example00`,
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
	testcases := []struct {
		name     string
		userId   string
		password string
	}{
		//各々テストケースを記述
		{
			name:     "example-00",
			userId:   `example.00@email.com`,
			password: `example00`,
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
