package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-malproxy/server/params"
)

/*
GetHtmlFromHostname => arg:hostname string/return:no return
					=> 指定のURLからHTMLファイルを取得する関数
*/
func GetHtmlFromHostname(hostname string) {
	url := params.UrlLinkedWithHost[hostname] // ホスト名からUPLを受け取る
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

	// fileName := hostname + ".html" // HTMLファイルの名前は、どのサイト由来か分かるように「`hostname`.html」で定義
	fileName := "index.html"

	err = ioutil.WriteFile(fileName, write, 0644) //ファイルに出力する
	if err != nil {
		panic(err)
	}

	// cmd := exec.Command("mv", fileName, "../../../")
	// if err = cmd.Run(); err != nil {
	// 	panic(err)
	// }
}
