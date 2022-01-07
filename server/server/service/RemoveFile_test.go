package service

import (
	"log"
	"net/url"
	"testing"
)

var testUrl = map[int]string{
	1: "https://www.google.com/",
	2: "https://www.amazon.co.jp/-/en/",
}

func TestRemoveFile(t *testing.T) {
	for _, Url := range testUrl {
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
