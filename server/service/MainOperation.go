package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func MainOperation(URL string) (string, error) {
	fmt.Println("\x1b[34mCatch:\x1b[0m", URL)
	//1. ルールに従って、URLを正規のものに戻す
	NewURL := strings.Replace(URL, "https://mitm.es3/", "https://", -1)
	//2. 正規URLで正規サーバーにアクセスし、返ってきたデータをHTMLファイルにして出力
	DecodeURL, err := UrlDecode(NewURL)
	if err != nil {
		return "err", nil
	}
	err = DataExtraction(DecodeURL)
	if err != nil {
		log.Fatal(err)
		return "err", err
	}
	//3. ルールに従って、URLを偽物のものに戻し、HTMLファイルとしてユーザーに返却
	u, err := url.Parse(DecodeURL)
	if err != nil {
		log.Fatal(err)
		return "err", err
	}
	fqdn := u.Hostname()
	err = ReadDataAndRewiteURL(fqdn)
	if err != nil {
		log.Fatal(err)
		return "err", err
	}
	return "rewrite_" + fqdn, nil
}

/*
ReadDataAndRewiteURL関数

FQDNを引数にとって、FQDN.htmlファイル内にある"https://"を"http://mitm.es3/"に置き換え、その結果を文字列として返す。
*/
func ReadDataAndRewiteURL(fqdn string) error {
	data, err := ioutil.ReadFile("/home/kimura/go-malproxy/server/templates/" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		log.Fatal(err)
		return err
	}
	res := `{{define "rewrite_` + fqdn + `"}}` + string(data) + `{{end}}` //データを文字列に変換
	Url, err := ExtractURL(res)
	if err != nil {
		return err
	}
	for i := range Url {
		fmt.Printf("\x1b[31mResult:%d = \x1b[0m%s", i, Url[i])
		fmt.Println()
	}
	rewrite := strings.Replace(res, `<a href="`, `<a href="http://localhost:3000/template?url=`, -1) //文字列の置き換え
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/rewrite_"+fqdn+".html", []byte(rewrite), os.ModePerm)
	// err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/rewrite_"+fqdn+".html", []byte(res), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return err
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
		return err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body) // []byte でリクエストの中身を取得
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/"+fileName+".html", byteArray, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

/*
ExtractURL関数

aタグにあるURLの抽出
*/
func ExtractURL(input string) ([]string, error) {
	aTag, err := regexp.Compile(`<a href="https://(.*)" `)
	if err != nil {
		return nil, err
	}
	link := aTag.FindAllString(input, -1) //aタグ内のハイパーリンクの抽出
	result := []string{}                  //結果の配列
	for i := 0; i < len(link); i++ {      //探索
		if strings.Contains(link[i], "&") {
			arr := []string{}
			arr = strings.Split(link[i], `"`)
			for _, v := range arr {
				isUrl, err := regexp.Compile(`https://(.*)`)
				if err != nil {
					return nil, err
				}
				if isUrl.Match([]byte(v)) {
					result = append(result, v)
				}
			}
		}
	}
	return result, nil
}

/*
UrlEncode関数

URLエンコーディング
*/
func UrlEncode(str string) (string, error) {
	str = url.QueryEscape(str)
	fmt.Println("\x1b[31mEncode:\x1b[0m", str)
	return str, nil
}

/*
UrlDecode関数

URLデコーディング
*/
func UrlDecode(str string) (string, error) {
	str, err := url.QueryUnescape(str)
	if err != nil {
		return "", err
	}
	fmt.Println("\x1b[31mDecode:\x1b[0m", str)
	return str, nil
}
