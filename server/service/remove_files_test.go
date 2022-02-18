package service

import (
	"log"
	"testing"
)

var RemoveList = []string{
	"https://www.google.com/",
	"https://www.amazon.co.jp/-/en/",
}

func TestRemoveFile(t *testing.T) {
	err := RemoveFile()
	if err != nil {
		log.Fatal(err)
	}
}
