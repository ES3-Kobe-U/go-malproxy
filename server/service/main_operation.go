package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/go-malproxy/server/params"
)

func MainOperation(URL string) (string, error) {
	//1. ルールに従って、URLを正規のものに戻す
	if judge := strings.Contains(URL, `https://mitm.es3/`); judge {
		URL = strings.Replace(URL, `https://mitm.es3/`, `https://`, -1)
	}
	if judge := strings.Contains(URL, `ANDANDAND`); judge {
		URL = strings.Replace(URL, `ANDANDAND`, `&`, -1)
	}
	if judge := strings.Contains(URL, `EQUALEQUALEQUAL`); judge {
		URL = strings.Replace(URL, `EQUALEQUALEQUAL`, `=`, -1)
	}

	//2. 抽出したURLをDecodeし，正規のURLとしてDecodedURLを得る．
	DecodedURL, err := UrlDecode(URL)
	if err != nil {
		return "err", nil
	}

	//3. 正規URL(DecodedURL)で正規サーバーにアクセスし、返ってきたデータをHTMLファイルにして出力
	err = DataExtraction(DecodedURL)
	if err != nil {
		log.Fatal(err)
		return "err", err
	}

	//4. ルールに従って、URLを偽物のものに戻し、HTMLファイルとしてユーザーに返却
	u, err := url.Parse(DecodedURL)
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
	return "autogen_rewrite_" + fqdn, nil
}

/*
ReadDataAndRewiteURL関数

FQDNを引数にとって、FQDN.htmlファイル内にある"https://"を"http://mitm.es3/"に置き換え、その結果を文字列として返す。
*/
func ReadDataAndRewiteURL(fqdn string) error {
	data, err := ioutil.ReadFile("server/templates/autogen_" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		log.Fatal(err)
		return err
	}
	res := `{{define "autogen_rewrite_` + fqdn + `"}}` + string(data) + `{{end}}` //データを文字列に変換
	Url, err := ExtractURL(res)
	if err != nil {
		return err
	}
	rew := strings.Replace(res, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	rew = strings.Replace(rew, `<a href='`, `<a href='/template?url=`, -1)  //文字列の置き換え
	rew = strings.Replace(rew, `pa3.min.js`, ``, -1)                        //楽天のCORSを回避する為に削除
	rew = strings.Replace(rew, params.RakutenLoginCode, params.ReplaceRakutenLoginCode, -1)
	for i := range Url {
		if judge := strings.Contains(rew, Url[i]); judge {
			Former := Url[i]
			Url[i] = strings.Replace(Url[i], "&amp;", "ANDANDAND", -1)
			Url[i] = strings.Replace(Url[i], "&", "ANDANDAND", -1)
			Url[i] = strings.Replace(Url[i], "%26", "ANDANDAND", -1)
			Url[i] = strings.Replace(Url[i], "=", "EQUALEQUALEQUAL", -1)
			rew = strings.Replace(rew, Former, Url[i], -1)
		}
	}
	rew = strings.Replace(rew, `&amp;`, `ANDANDAND`, -1)
	rew = strings.Replace(rew, `%26`, `ANDANDAND`, -1)
	err = ioutil.WriteFile("server/templates/autogen_rewrite_"+fqdn+".html", []byte(rew), os.ModePerm)
	if err != nil {
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
	err = ioutil.WriteFile("server/templates/autogen_"+fileName+".html", byteArray, os.ModePerm)
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
			arr := strings.Split(link[i], `"`)
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
	aTag, err = regexp.Compile(`<a href='https://(.*)" `)
	if err != nil {
		return nil, err
	}
	link = aTag.FindAllString(input, -1) //aタグ内のハイパーリンクの抽出
	for i := 0; i < len(link); i++ {     //探索
		if strings.Contains(link[i], "&") {
			arr := strings.Split(link[i], `'`)
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
	return str, nil
}

/*
UrlDecode関数

URLデコーディング
*/
func UrlDecode(str string) (string, error) {
	_, err := url.QueryUnescape(str)
	if err != nil {
		return "", err
	}
	return str, nil
}
