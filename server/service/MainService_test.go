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
	"testing"
)

var ListOfURL = []string{
	"https://mitm.es3/amazon.co.jp/",
	"https://mitm.es3/github.com/marketplace/circleci",
	"https://mitm.es3/hub.docker.com/search?type=image",
	"https://mitm.es3/qiita.com/official-events/5cb794f7cb9ac194ed70",
}

func TestMainService(t *testing.T) {
	for i := range ListOfURL {
		err := MainService(ListOfURL[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

// var GoogleSearch_Sample1 = []string{
// 	"神戸大学",
// 	"神戸 大学",
// 	"神戸　大学",
// 	"KobeUniv.",
// 	"Kobe Univ.",
// 	"Kobe　Univ.",
// }

var GoogleSearch_Sample1 = []string{
	"Amazon 通販",
}

var GoogleSearch_Sample2 = []string{
	"神戸大学",
	"神戸 大学",
	"神戸　大学",
	"KobeUniv.",
	"Kobe Univ.",
	"Kobe　Univ.",
	"神戸大学",
	"神戸  大学",
	"神戸　　大学",
	"KobeUniv.",
	"Kobe  Univ.",
	"Kobe　　Univ.",
}

func TestGoogleSearch(t *testing.T) {
	for _, word := range GoogleSearch_Sample1 {
		err := GoogleSearch(word)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestMakeURL(t *testing.T) {
	for _, word := range GoogleSearch_Sample2 {
		var query string
		cnt := 0
		for _, char := range word {
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
		fmt.Println(URL)
	}
}

var ReadData_Sample = []string{
	"www.google.com",
	//"www.amazon.co.jp",
}

func TestReadDataAndRewiteURL(t *testing.T) {
	for _, fqdn := range ReadData_Sample {
		res, err := ReadDataAndRewiteURL(fqdn)
		if err != nil {
			log.Fatal(err)
		}
		//ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/rewite_"+fqdn+".html", []byte(res), os.ModePerm)
		ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/index.html", []byte(res), os.ModePerm)
	}
}

var RemoveFile_Sample1 = []string{
	"https://www.google.com/",
	"https://www.amazon.co.jp/-/en/",
}

func TestRemoveFile(t *testing.T) {
	for _, Url := range RemoveFile_Sample1 {
		u, err := url.Parse(Url)
		if err != nil {
			log.Fatal(err)
		}
		err = RemoveFile(u.Hostname())
		if err != nil {
			log.Fatal(err)
		}
	}
}

var DataExtraction_Sample = []string{
	//"https://www.google.com/",
	"https://www.amazon.co.jp/-/en/",
	//"https://rakuten.co.jp",
}

func TestDataExtraction(t *testing.T) {
	for _, Url := range DataExtraction_Sample {
		resp, _ := http.Get(Url)
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)
		u, err := url.Parse(Url)
		if err != nil {
			log.Fatal(err)
		}
		fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
		err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/"+fileName+".html", byteArray, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestGetURL(t *testing.T) {
	// url の指定
	url := "https://rakuten.co.jp/"
	// 正規表現の作成
	re, err := regexp.Compile("http(.*)://(.*)")
	if err != nil {
		return
	}
	// net/http でのリクエストの発射
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	// []byte でリクエストの中身を取得
	byteArray, _ := ioutil.ReadAll(resp.Body)
	// 正規表現にあったものを全てlinks に入れる
	links := re.FindAllString(string(byteArray), -1)
	for i := 0; i < len(links); i++ {
		fmt.Println(links[i])
	}
}

/*
URLの中身を取得　⇒　中身からハイパーリンクのURLを取得　⇒　URLを書き換え　⇒　htmlに出力
*/
func TestGetURLAndOutputHtml(t *testing.T) {
	Url := "https://amazon.co.jp/"
	re, err := regexp.Compile("http(.*)://(.*)")
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := http.Get(Url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	links := re.FindAllString(string(byteArray), -1)
	m := map[string]string{}
	for i := 0; i < len(links); i++ {
		u, err := url.Parse(links[i])
		if err != nil {
			log.Fatal(err)
		}
		fn := u.Hostname()
		fmt.Println("\x1b[31mHost Name ---> \x1b[0m", fn)
		m["https://"+fn] = "http://mitm.es3/" + fn
		rewrite := strings.Replace(links[i], "https://"+fn, "http://mitm.es3/"+fn, -1)
		fmt.Println("\x1b[31mRewrite ---> \x1b[0m", rewrite)
	}
	for i, j := range m {
		fmt.Println(i)
		fmt.Println(j)
	}
	u, err := url.Parse(Url)
	if err != nil {
		log.Fatal(err)
	}
	fileName := u.Hostname()
	err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/"+fileName+".html", byteArray, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
