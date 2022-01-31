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

var ToEncodeUrlList = []string{
	"https://www.amazon.co.jp/ap/signin?openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.co.jp%2F%3F_encoding%3DUTF8%26ref_%3Dnav_ya_signin&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.assoc_handle=jpflex&openid.mode=checkid_setup&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&",
	"https://www.amazon.co.jp/ap/signin?openid.mode=checkid_setup&amp;openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&amp;openid.return_to=https%3A%2F%2Fwww.amazon.co.jp%2Fref%3Dgw_sgn_ib%3F_encoding%3DUTF8%26pf_rd_p%3Df1e47293-143b-4dd7-9f05-590bea73bcbe%26pd_rd_wg%3D9B9dc%26pf_rd_r%3D01VZFGBVHH03P17J2DWJ%26pd_rd_w%3DY24yn%26pd_rd_r%3Dd7d8522c-13e5-4adc-b67d-4c0ceeb1f97d&amp;openid.assoc_handle=jpflex&amp;openid.pape.max_auth_age=0",
}

var ToDecodeUrlList = []string{
	"https%3A%2F%2Fwww.amazon.co.jp%2Fap%2Fsignin%3Fopenid.pape.max_auth_age%3D0%26openid.return_to%3Dhttps%253A%252F%252Fwww.amazon.co.jp%252F%253F_encoding%253DUTF8%2526ref_%253Dnav_ya_signin%26openid.identity%3Dhttp%253A%252F%252Fspecs.openid.net%252Fauth%252F2.0%252Fidentifier_select%26openid.assoc_handle%3Djpflex%26openid.mode%3Dcheckid_setup%26openid.claimed_id%3Dhttp%253A%252F%252Fspecs.openid.net%252Fauth%252F2.0%252Fidentifier_select%26openid.ns%3Dhttp%253A%252F%252Fspecs.openid.net%252Fauth%252F2.0%26",
	"https%3A%2F%2Fwww.amazon.co.jp%2Fap%2Fsignin%3Fopenid.mode%3Dcheckid_setup%26amp%3Bopenid.ns%3Dhttp%253A%252F%252Fspecs.openid.net%252Fauth%252F2.0%26amp%3Bopenid.return_to%3Dhttps%253A%252F%252Fwww.amazon.co.jp%252Fref%253Dgw_sgn_ib%253F_encoding%253DUTF8%2526pf_rd_p%253Df1e47293-143b-4dd7-9f05-590bea73bcbe%2526pd_rd_wg%253D9B9dc%2526pf_rd_r%253D01VZFGBVHH03P17J2DWJ%2526pd_rd_w%253DY24yn%2526pd_rd_r%253Dd7d8522c-13e5-4adc-b67d-4c0ceeb1f97d%26amp%3Bopenid.assoc_handle%3Djpflex%26amp%3Bopenid.pape.max_auth_age%3D0",
}

func TestMainOperation(t *testing.T) {
	for i := range DummyUrlList {
		res, err := MainOperation(DummyUrlList[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("FQDN:", res)
	}
}

func TestReadDataAndRewiteURL(t *testing.T) {
	for _, fqdn := range UrlList {
		err := ReadDataAndRewiteURL(fqdn)
		if err != nil {
			log.Fatal(err)
		}
		//ioutil.WriteFile("/home/kimura/go-malproxy/server/templates/index.html", []byte(res), os.ModePerm)
	}
}

func TestDataExtraction(t *testing.T) {
	for _, Url := range UrlList {
		resp, err := http.Get(Url)
		if err != nil {
			log.Fatal(err)
		}
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
		resp, err := http.Get(Url) // net/http でのリクエストの発射
		if err != nil {
			log.Fatal(err)
		}
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
		resp, err := http.Get(Url)
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
