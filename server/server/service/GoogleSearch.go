package service

func GoogleSearch(Word string) error {
	var search string
	for _, char := range Word {
		if char == ' ' || char == '　' { //空文字列なら+に変換
			char = '+'
		}
		search = search + string(char)
	}
	URL := "https://google.com/search?q=" + search //Google検索のURLはこれで統一されているっぽい
	err := DataExtraction(URL)
	if err != nil {
		return err
	}
	return nil
}
