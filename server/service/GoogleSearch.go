package service

// Google検索用のURLの作成処理
// func GoogleSearch(Word string) error {
// 	var query string
// 	cnt := 0
// 	for _, char := range Word {
// 		if char == ' ' || char == '　' { //空文字列なら+に変換
// 			cnt += 1
// 			if cnt < 2 {
// 				char = '+'
// 				query = query + string(char)
// 			}
// 		} else {
// 			cnt = 0
// 			query = query + string(char)
// 		}
// 	}
// 	URL := "https://google.com/search?q=" + query //Google検索のURLはこれで統一されているっぽい
// 	err := DataExtraction(URL)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
