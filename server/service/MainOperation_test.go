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

var UrlList = []string{
	"https://amazon.co.jp/",
	"https://github.com/marketplace/circleci",
	"https://hub.docker.com/search?type=image",
	"https://qiita.com/official-events/5cb794f7cb9ac194ed70",
}

var DummyUrlList = []string{
	"https://mitm.es3/amazon.co.jp/",
	"https://mitm.es3/github.com/marketplace/circleci",
	"https://mitm.es3/hub.docker.com/search?type=image",
	"https://mitm.es3/qiita.com/official-events/5cb794f7cb9ac194ed70",
}

func TestMainOperation(t *testing.T) {
	for i := range DummyUrlList {
		err := MainOperation(DummyUrlList[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestReadDataAndRewiteURL(t *testing.T) {
	for _, fqdn := range UrlList {
		res, err := ReadDataAndRewiteURL(fqdn)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/index.html", []byte(res), os.ModePerm)
	}
}

func TestRemoveFile(t *testing.T) {
	for _, Url := range UrlList {
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

func TestDataExtraction(t *testing.T) {
	for _, Url := range UrlList {
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
	for _, Url := range UrlList {
		re, err := regexp.Compile("https://(.*)") // 正規表現の作成
		if err != nil {
			return
		}
		resp, _ := http.Get(Url) // net/http でのリクエストの発射
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)        // []byte でリクエストの中身を取得
		links := re.FindAllString(string(byteArray), -1) // 正規表現にあったものを全てlinks に入れる
		for i := 0; i < len(links); i++ {
			fmt.Println("\x1b[31m取得結果:\x1b[0m", links[i])
		}
	}
}

/*
URLの中身を取得　⇒　中身からハイパーリンクのURLを取得　⇒　URLを書き換え　⇒　htmlに出力
*/
func TestGetURLAndOutputHtml(t *testing.T) {
	for _, Url := range UrlList {
		re, err := regexp.Compile("https://(.*)")
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
}
