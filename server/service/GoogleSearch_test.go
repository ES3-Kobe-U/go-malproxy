package service

import (
	"fmt"
	"log"
	"testing"
)

var GoogleSearch_Sample1 = []string{
	"神戸大学",
	"神戸 大学",
	"神戸　大学",
	"KobeUniv.",
	"Kobe Univ.",
	"Kobe　Univ.",
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

func TestMakeGoogleURL(t *testing.T) {
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
