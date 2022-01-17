package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func MainService(URL string) error {
	//1. 叩いたURLの取得
	//2. ルールに従って、URLを正規のものに戻す
	//3. 正規URLで正規サーバーにアクセスし、帰ってきたデータをHTMLファイルにして出力
	//4. ルールに従って、URLを偽物のものに戻す
	//5. 偽物のURLに変換したデータをHTMLファイルとしてユーザーに返す。
	return nil
}

// Google検索用のURLの作成処理
func GoogleSearch(Word string) error {
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
	URL := "https://www.google.com/search?q=" + query //Google検索のURLはこれで統一されているっぽい
	err := DataExtraction(URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//データを読み込み&URLの書き換え
func ReadDataAndRewiteURL(fqdn string) (string, error) {
	data, err := ioutil.ReadFile("test/" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		return "", err
	}
	res := string(data)                                                    //データを文字列に変換
	rewrite := strings.Replace(res, "https://", "http://go-malproxy/", -1) //文字列の置き換え
	return rewrite, nil
}

//作成したHTMLファイルの削除
func RemoveFile(fqdn string) error {
	cmdRemove := exec.Command("rm", "test/"+fqdn+".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	cmdRemove.Stderr = os.Stderr
	cmdRemove.Stdin = os.Stdin
	out, err := cmdRemove.Output()
	if err != nil {
		log.Println("Err", err)
		return err
	} else {
		log.Println("remove files --> success:", string(out))
	}
	return nil
}

//URLからHTMLファイルの取得&自動生成
func DataExtraction(URL string) error {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		fmt.Print("url scarapping failed")
		return err
	}
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, err := doc.Html()
	if err != nil {
		fmt.Print("dom get failed")
		return err
	}
	fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/server/service/test/"+fileName+".html", []byte(res), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
