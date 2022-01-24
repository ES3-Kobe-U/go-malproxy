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

func MainService(URL string) error {
	//1. 叩いたURLの取得
	//2. ルールに従って、URLを正規のものに戻す
	//3. 正規URLで正規サーバーにアクセスし、帰ってきたデータをHTMLファイルにして出力
	//4. ルールに従って、URLを偽物のものに戻す
	//5. 偽物のURLに変換したデータをHTMLファイルとしてユーザーに返す。
	return nil
}

/*
GoogleSearch関数

検索ワードを引数にとって、その検索結果のhtmlファイルを自動生成する。
*/
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

/*
ReadDataAndRewiteURL関数

FQDNを引数にとって、FQDN.htmlファイル内にある"https://"を"http://go-malproxy/"に置き換え、その結果を文字列として返す。
*/
func ReadDataAndRewiteURL(fqdn string) (string, error) {
	data, err := ioutil.ReadFile("/home/kimura/go-malproxy/server/server/views/" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		return "", err
	}
	res := string(data) //データを文字列に変換
	//m := map[string]int{}                                                          //urlのパターンをマップで管理
	r, err := regexp.Compile("https://(.*)")
	if err != nil {
		log.Fatal(err)
	}
	links := r.FindAllString(res, -1)
	for i := 0; i < len(links); i++ {
		fmt.Println(links[i])
	}
	rewrite := strings.Replace(res, "https://(*)/", "http://localhost:1323/", -1) //文字列の置き換え
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
	// doc, err := goquery.NewDocument(URL)
	// if err != nil {
	// 	fmt.Print("url scarapping failed")
	// 	return err
	// }
	// url の指定
	re, err := regexp.Compile("http(.*)://(.*)")
	if err != nil {
		return err
	}
	// net/http でのリクエストの発射
	resp, _ := http.Get(URL)
	defer resp.Body.Close()
	// []byte でリクエストの中身を取得
	byteArray, _ := ioutil.ReadAll(resp.Body)
	// 正規表現にあったものを全てlinks に入れる
	links := re.FindAllString(string(byteArray), -1)
	for i := 0; i < len(links); i++ {
		fmt.Println(links[i])
	}
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// res, err := doc.Html()
	// if err != nil {
	// 	fmt.Print("dom get failed")
	// 	return err
	// }
	fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
	// err = ioutil.WriteFile("/home/kimura/go-malproxy/server/server/views/"+fileName+".html", []byte(res), os.ModePerm)
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/server/views/"+fileName+".html", byteArray, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
