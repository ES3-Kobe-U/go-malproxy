package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type DataExtractionService interface {
	DataExtraction(URL string)
}

//URLからHTMLファイルを取得&自動生成
func DataExtraction(URL string) {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	res, err := doc.Html()
	if err != nil {
		fmt.Print("dom get failed")
	}
	fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
	ioutil.WriteFile("../../../client/base/src/test/"+fileName+".html", []byte(res), os.ModePerm)
}
