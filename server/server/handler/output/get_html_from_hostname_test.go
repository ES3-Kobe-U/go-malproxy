package handler

import (
	"fmt"
	"testing"
)

var HtmlFromHostnameTests = map[string]string{
	"rakuten.co.jp": "http://rakuten.co.jp/",
	"golang.org":    "http://golang.org/",
}

func TestGetHtmlFromHostname(t *testing.T) {
	for hostname := range HtmlFromHostnameTests {
		fmt.Println("src = ", hostname)
		GetHtmlFromHostname(hostname)
	}
}
