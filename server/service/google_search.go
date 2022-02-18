package service

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

/*
GoogleSearch関数

検索ワードを引数にとって、その検索結果のhtmlファイルを自動生成する。
*/
func GoogleSearch(Word string) (string, error) {
	var query string
	cnt := 0
	for _, char := range Word {
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
	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	fqdn := u.Hostname()
	err = DataExtraction(URL)
	if err != nil {
		return "", err
	}
	return fqdn, nil
}

func RewriteUrlOfGoogleSearch(fqdn string) error {
	data, err := ioutil.ReadFile("/home/kimura/go-malproxy/server/templates/autogen_" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		log.Fatal(err)
		return err
	}
	res := `{{define "autogen_rewrite_` + fqdn + `"}}` + string(data) + `{{end}}`                           //データを文字列に変換
	rewrite := strings.Replace(res, `<a href="/url?q=`, `<a href="http://localhost:3333/template?url=`, -1) //文字列の置き換え
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/autogen_rewrite_"+fqdn+".html", []byte(rewrite), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
