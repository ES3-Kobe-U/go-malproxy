package service

import (
	"log"
	"testing"
)

var Word = map[int]string{
	1: "神戸大学",
	2: "神戸 大学",
	3: "神戸　大学",
	4: "KobeUniv.",
	5: "Kobe Univ.",
	6: "Kobe　Univ.",
}

func TestGoogleSearch(t *testing.T) {
	for _, word := range Word {
		err := GoogleSearch(word)
		if err != nil {
			log.Fatal(err)
		}
	}
}
