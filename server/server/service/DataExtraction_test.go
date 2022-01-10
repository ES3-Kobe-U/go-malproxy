package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var URL = map[int]string{
	1: "https://www.google.com/",
	2: "https://www.amazon.co.jp/-/en/",
	3: "https://www.google.com/search?q=Golang",
}

func TestDataExtraction(t *testing.T) {
	for _, Url := range URL {
		doc, err := goquery.NewDocument(Url)
		if err != nil {
			fmt.Print("url scarapping failed")
		}
		u, err := url.Parse(Url)
		if err != nil {
			log.Fatal(err)
		}
		res, err := doc.Html()
		if err != nil {
			fmt.Print("dom get failed")
		}
		fileName := u.Hostname()
		ioutil.WriteFile("test/"+fileName+".html", []byte(res), os.ModePerm)
	}
}
