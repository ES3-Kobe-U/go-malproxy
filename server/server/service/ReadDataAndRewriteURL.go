package service

import (
	"io/ioutil"
	"strings"
)

type ReadDataAndRewiteURLService interface {
	ReadDataAndRewiteURL(fqdn string) (string, error)
}

func ReadDataAndRewiteURL(fqdn string) (string, error) {
	data, err := ioutil.ReadFile("test/" + fqdn + ".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
	if err != nil {
		return "", err
	}
	res := string(data)                                                        //データを文字列に変換
	rewrite := strings.Replace(res, "https://"+fqdn, "http://go-malproxy", -1) //文字列の置き換え
	return rewrite, nil
}
