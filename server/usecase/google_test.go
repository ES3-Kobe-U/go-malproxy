package usecase

import (
	"fmt"
	"log"
	"testing"
)

func TestGoogleSearch(t *testing.T) {
	testcases := []struct {
		name string
		word string
	}{
		{
			name: "example_00",
			word: "神戸大学",
		},
		{
			name: "example_01",
			word: "神戸 大学",
		},
		{
			name: "example_02",
			word: "神戸　大学",
		},
		{
			name: "example_03",
			word: "KobeUniv.",
		},
		{
			name: "example_04",
			word: "Kobe Univ.",
		},
		{
			name: "example_05",
			word: "Kobe　Univ.",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := GoogleSearch(testcase.word)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("res:", res)
		})
	}
}
func TestMakeGoogleURL(t *testing.T) {
	testcases := []struct {
		name   string
		word   string
		expect string
	}{
		{
			name:   "example_00",
			word:   "神戸大学",
			expect: "https://google.com/search?q=神戸大学",
		},
		{
			name:   "example_01",
			word:   "神戸 大学",
			expect: "https://google.com/search?q=神戸+大学",
		},
		{
			name:   "example_02",
			word:   "神戸　大学",
			expect: "https://google.com/search?q=神戸+大学",
		},
		{
			name:   "example_03",
			word:   "KobeUniv.",
			expect: "https://google.com/search?q=KobeUniv.",
		},
		{
			name:   "example_04",
			word:   "Kobe Univ.",
			expect: "https://google.com/search?q=Kobe+Univ.",
		},
		{
			name:   "example_05",
			word:   "Kobe　Univ.",
			expect: "https://google.com/search?q=Kobe+Univ.",
		},
		{
			name:   "example_06",
			word:   "神戸  大学",
			expect: "https://google.com/search?q=神戸+大学",
		},
		{
			name:   "example_07",
			word:   "神戸　　大学",
			expect: "https://google.com/search?q=神戸+大学",
		},
		{
			name:   "example_08",
			word:   "Kobe  Univ.",
			expect: "https://google.com/search?q=Kobe+Univ.",
		},
		{
			name:   "example_09",
			word:   "Kobe　　Univ.",
			expect: "https://google.com/search?q=Kobe+Univ.",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			var query string
			cnt := 0
			for _, char := range testcase.word {
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
			if URL != testcase.expect {
				t.Errorf("got=%s, want=%s\n", URL, testcase.expect)
			}
		})
	}
}
