package service

import (
	"log"
	"testing"
)

var FullDownloadOfTheWebPage_Sample = []string{
	"https://www.kobe-u.ac.jp/",
}

func TestFullDownloadOfTheWebPage(t *testing.T) {
	for _, Url := range FullDownloadOfTheWebPage_Sample {
		err := FullDownloadOfTheWebPage(Url)
		if err != nil {
			log.Fatal(err)
		}
	}
}