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

func TestMainService(t *testing.T) {

}

var GoogleSearch_Sample1 = []string{
	"神戸大学",
	"神戸 大学",
	"神戸　大学",
	"KobeUniv.",
	"Kobe Univ.",
	"Kobe　Univ.",
}

var GoogleSearch_Sample2 = []string{
	"神戸大学",
	"神戸 大学",
	"神戸　大学",
	"KobeUniv.",
	"Kobe Univ.",
	"Kobe　Univ.",
	"神戸大学",
	"神戸  大学",
	"神戸　　大学",
	"KobeUniv.",
	"Kobe  Univ.",
	"Kobe　　Univ.",
}

func TestGoogleSearch(t *testing.T) {
	for _, word := range GoogleSearch_Sample1 {
		err := GoogleSearch(word)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestMakeURL(t *testing.T) {
	for _, word := range GoogleSearch_Sample2 {
		var query string
		cnt := 0
		for _, char := range word {
			if char == ' ' || char == '　' { //空文字列なら+に変換
				cnt += 1
				if cnt < 2 {
					char = '+'
					query = query + string(char)
				}
			} else {
				cnt = 0
				query = query + string(char)
			}
		}
		URL := "https://google.com/search?q=" + query //Google検索のURLはこれで統一されているっぽい
		fmt.Println(URL)
	}
}

var ReadData_Sample = []string{
	"www.google.com",
	"www.amazon.co.jp",
}

func TestReadData(t *testing.T) {
	for _, fqdn := range ReadData_Sample {
		res, err := ReadDataAndRewiteURL(fqdn)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile("test/rewite_"+fqdn+".html", []byte(res), os.ModePerm)
	}
}

var RemoveFile_Sample1 = []string{
	"https://www.google.com/",
	"https://www.amazon.co.jp/-/en/",
}

func TestRemoveFile(t *testing.T) {
	for _, Url := range RemoveFile_Sample1 {
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

var DataExtraction_Sample = []string{
	"https://www.google.com/",
	"https://www.amazon.co.jp/-/en/",
	"https://www.google.com/search?q=Golang",
}

func TestDataExtraction(t *testing.T) {
	for _, Url := range DataExtraction_Sample {
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
