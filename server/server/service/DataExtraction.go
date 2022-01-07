package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

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

	fileName := u.Hostname()

	ioutil.WriteFile("../../../client/base/src/test/"+fileName+".html", []byte(res), os.ModePerm)
}
