package service

import "net/url"

/*
GoogleSearch関数

検索ワードを引数にとって、その検索結果のhtmlファイルを自動生成する。
*/
func GoogleSearch(Word string) (string, error) {
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
	URL := "https://google.com/search?q=" + query //Google検索のURLはこれで統一されているっぽい
	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	fqdn := u.Hostname()
	err = DataExtraction(URL)
	if err != nil {
		return "", err
	}
	return fqdn, nil
}
