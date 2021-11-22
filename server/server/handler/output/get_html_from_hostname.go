package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ホスト名と表示する画面をマップで管理しておく
// TODO: 後々、各サイトのログイン画面のURLに変更する
var HtmlFromHostname = map[string]string{
	"rakuten.co.jp": "http://rakuten.co.jp/",
	"golang.org":    "http://golang.org/",
}

func GetHtmlFromHostname(hostname string) {
	url := HtmlFromHostname[hostname] // ホスト名からUPLを受け取る
	fmt.Printf("HTML code of %s ...\n", url)

	resp, err := http.Get(url) //レスポンスを受け取る
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() //レスポンスを閉じる

	html, err := ioutil.ReadAll(resp.Body) // htmlファイルをbyteのスライス形式で受け取る
	if err != nil {
		panic(err)
	}

	var write []byte
	write = append(write, html...) //出力結果を入れていく

	fileName := hostname + ".html" // HTMLファイルの名前は、どのサイト由来か分かるように「`hostname`.html」で定義

	err = ioutil.WriteFile(fileName, write, 0644) //ファイルに出力する
	if err != nil {
		panic(err)
	}
}
