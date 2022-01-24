package service

import (
	"log"
	"net/url"
	"testing"
)

var RemoveList = []string{
	"https://www.google.com/",
	"https://www.amazon.co.jp/-/en/",
}

func TestRemoveFile(t *testing.T) {
	for _, Url := range RemoveList {
		u, err := url.Parse(Url)
		if err != nil {
			log.Fatal(err)
		}
		err = RemoveFile(u.Hostname())
		if err != nil {
			log.Fatal(err)
		}
	}
}
