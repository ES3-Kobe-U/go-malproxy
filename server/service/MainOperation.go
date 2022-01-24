package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func MainOperation(URL string) error {
	//1. ルールに従って、URLを正規のものに戻す
	NewURL := strings.Replace(URL, "https://mitm.es3/", "https://", -1)

	//2. 正規URLで正規サーバーにアクセスし、返ってきたデータをHTMLファイルにして出力
	err := DataExtraction(NewURL)
	if err != nil {
		return err
	}

	//3. ルールに従って、URLを偽物のものに戻し、HTMLファイルとしてユーザーに返却
	u, err := url.Parse(NewURL)
	if err != nil {
		return err
	}
	fqdn := u.Hostname()
	_, err = ReadDataAndRewiteURL(fqdn)
	if err != nil {
		return err
	}

	return nil
}

/*
ReadDataAndRewiteURL関数

FQDNを引数にとって、FQDN.htmlファイル内にある"https://"を"http://mitm.es3/"に置き換え、その結果を文字列として返す。
*/
func ReadDataAndRewiteURL(fqdn string) (string, error) {
	data, err := ioutil.ReadFile("/home/kimura/go-malproxy/server/templates/" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		return "", err
	}
	res := string(data) //データを文字列に変換
	r, err := regexp.Compile("https://(.*)")
	if err != nil {
		log.Fatal(err)
	}
	links := r.FindAllString(res, -1)
	for i := 0; i < len(links); i++ {
		fmt.Println("\x1b[34m 抜き出したURL---> \x1b[0m", links[i])
	}
	rewrite := strings.Replace(res, "https://", "https://mitm.es3/", -1) //文字列の置き換え
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/rewrite_"+fqdn+".html", []byte(rewrite), os.ModePerm)
	if err != nil {
		return "", err
	}
	fmt.Println("\x1b[35m 書き換え結果---> \x1b[0m", rewrite)
	return rewrite, nil
}

/*
RemoveFile関数

FQDNを引数にとって、FQDN.htmlを外部コマンドで削除する。
*/
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

/*
DataExtraction関数

実際の正規URLを引数にとって、htmlファイルを自動生成する。
*/
func DataExtraction(URL string) error {
	resp, err := http.Get(URL) // net/http でのリクエストの発射
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body) // []byte でリクエストの中身を取得

	re, err := regexp.Compile("https://(.*)") // パターンの指定
	if err != nil {
		return err
	}
	links := re.FindAllString(string(byteArray), -1) // 正規表現にあったものを全てlinks に入れる
	for i := 0; i < len(links); i++ {
		fmt.Println("\x1b[31m 正規表現---> \x1b[0m", links[i])
	}

	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/"+fileName+".html", byteArray, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
