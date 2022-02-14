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

func TestMainOperation(t *testing.T) {
	testcases := []struct {
		name string
		url  string
	}{
		{
			name: "example_00",
			url:  "https://mitm.es3/amazon.co.jp/",
		},
		{
			name: "example_01",
			url:  "https://mitm.es3/github.com/marketplace/circleci",
		},
		{
			name: "example_02",
			url:  "https://mitm.es3/hub.docker.com/search?type=image",
		},
		{
			name: "example_03",
			url:  "https://mitm.es3/qiita.com/official-events/5cb794f7cb9ac194ed70",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := MainOperation(testcase.url)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\x1b[35mFQDN:\x1b[0m", res)
		})
	}
}

func TestReadDataAndRewiteURL(t *testing.T) {
	testcases := []struct {
		name string
		url  string
	}{
		{
			name: "example_00",
			url:  "https://amazon.co.jp/",
		},
		{
			name: "example_01",
			url:  "https://github.com/marketplace/circleci",
		},
		{
			name: "example_02",
			url:  "https://hub.docker.com/search?type=image",
		},
		{
			name: "example_03",
			url:  "https://qiita.com/official-events/5cb794f7cb9ac194ed70",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			err := ReadDataAndRewiteURL(testcase.url)
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestDataExtraction(t *testing.T) {
	testcases := []struct {
		name string
		url  string
	}{
		{
			name: "example_00",
			url:  "https://amazon.co.jp/",
		},
		{
			name: "example_01",
			url:  "https://github.com/marketplace/circleci",
		},
		{
			name: "example_02",
			url:  "https://hub.docker.com/search?type=image",
		},
		{
			name: "example_03",
			url:  "https://qiita.com/official-events/5cb794f7cb9ac194ed70",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			resp, err := http.Get(testcase.url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			byteArray, _ := ioutil.ReadAll(resp.Body)
			u, err := url.Parse(testcase.url)
			if err != nil {
				log.Fatal(err)
			}
			fileName := u.Hostname() //ファイル名はホスト名で統一（多分FQDNの形で返されるので、以後変数名はfqdnで統一したい）
			err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/"+fileName+".html", byteArray, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestGetURL(t *testing.T) {
	testcases := []struct {
		name string
		url  string
	}{
		{
			name: "example_00",
			url:  "https://amazon.co.jp/",
		},
		{
			name: "example_01",
			url:  "https://github.com/marketplace/circleci",
		},
		{
			name: "example_02",
			url:  "https://hub.docker.com/search?type=image",
		},
		{
			name: "example_03",
			url:  "https://qiita.com/official-events/5cb794f7cb9ac194ed70",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			re, err := regexp.Compile("https://(.*)") // 正規表現の作成
			if err != nil {
				return
			}
			resp, err := http.Get(testcase.url) // net/http でのリクエストの発射
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			byteArray, _ := ioutil.ReadAll(resp.Body)        // []byte でリクエストの中身を取得
			links := re.FindAllString(string(byteArray), -1) // 正規表現にあったものを全てlinks に入れる
			for i := 0; i < len(links); i++ {
				fmt.Println("\x1b[31m取得結果:\x1b[0m", links[i])
			}
		})
	}
}

/*
URLの中身を取得　⇒　中身からハイパーリンクのURLを取得　⇒　URLを書き換え　⇒　htmlに出力
*/
func TestGetURLAndOutputHtml(t *testing.T) {
	testcases := []struct {
		name string
		url  string
	}{
		{
			name: "example_00",
			url:  "https://amazon.co.jp/",
		},
		{
			name: "example_01",
			url:  "https://github.com/marketplace/circleci",
		},
		{
			name: "example_02",
			url:  "https://hub.docker.com/search?type=image",
		},
		{
			name: "example_03",
			url:  "https://qiita.com/official-events/5cb794f7cb9ac194ed70",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			re, err := regexp.Compile("https://(.*)")
			if err != nil {
				log.Fatal(err)
			}
			resp, err := http.Get(testcase.url)
			if err != nil {
				log.Fatal(err)
			}
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
			u, err := url.Parse(testcase.url)
			if err != nil {
				log.Fatal(err)
			}
			fileName := u.Hostname()
			err = ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/"+fileName+".html", byteArray, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestExtractURL(t *testing.T) {
	data, err := ioutil.ReadFile("/home/kimura/go-malproxy/server/templates/www.amazon.co.jp.html")
	if err != nil {
		log.Fatal(err)
	}
	res := string(data)
	result, err := ExtractURL(res)
	if err != nil {
		log.Fatal(err)
	}
	for i := range result {
		fmt.Printf("\x1b[31mResult:%d = \x1b[0m%s", i, result[i])
		fmt.Println()
	}
}

func TestUrlEncodeAndDecode(t *testing.T) {
	data, err := ioutil.ReadFile("/home/kimura/go-malproxy/server/templates/www.amazon.co.jp.html")
	if err != nil {
		log.Fatal(err)
	}
	res := string(data)
	result, err := ExtractURL(res)
	if err != nil {
		log.Fatal(err)
	}
	for i := range result {
		fmt.Printf("\x1b[31mResult:%d = \x1b[0m%s", i, result[i])
		fmt.Println()
	}

	encodeRes := []string{}
	for _, str := range result {
		res, err := UrlEncode(str)
		if err != nil {
			log.Fatal(err)
		}
		encodeRes = append(encodeRes, res)
	}

	for _, str := range encodeRes {
		_, err := UrlDecode(str)
		if err != nil {
			log.Fatal(err)
		}
	}
}
