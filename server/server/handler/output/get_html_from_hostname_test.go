package handler

import (
	"fmt"
	"testing"
)

var TestCase []string = []string{
	//"rakuten.co.jp",
	//"golang.org",
	"amazon.co.jp",
}

func TestGetHtmlFromHostname(t *testing.T) {
	for _, hostname := range TestCase {
		fmt.Println("src = ", hostname)
		GetHtmlFromHostname(hostname)
	}
}
