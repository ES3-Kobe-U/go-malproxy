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
	"amazon.co.jp":  "http://www.amazon.co.jp/ap/signin?_encoding=UTF8&language=en&openid.assoc_handle=jpflex&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.ns.pape=http%3A%2F%2Fspecs.openid.net%2Fextensions%2Fpape%2F1.0&openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.co.jp%2Fgp%2Fyourstore%2Fhome%3Fie%3DUTF8%26action%3Dsign-out%26language%3Den%26path%3D%252Fgp%252Fyourstore%252Fhome%26ref_%3Dnav_AccountFlyout_signout%26signIn%3D1%26useRedirectOnSuccess%3D1",
}

/*
GetHtmlFromHostname => arg:hostname string/return:no return
					=> 指定のURLからHTMLファイルを取得する関数
*/
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
