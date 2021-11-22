package handler

import (
	"fmt"
	"testing"
)

var hostnameFromIPTests = map[string]string{
	"127.0.0.1": "localhost",
	"8.8.4.4":   "dns.google.",
	//"52.119.168.48": "amazon.co.jp",
}

func TestGetHostnameFromIPAddress(t *testing.T) {
	for ip, hn := range hostnameFromIPTests { // ipアドレスとホスト名の配列を探索
		op, err := GetHostnameFromIPAddress(ip) // ipアドレスからホスト名を得る
		if err != nil {
			t.Error(err)
		}
		judge := false // 判定
		var key string // 実行結果から得られたホスト名
		for _, key = range op {
			if key == hn {
				fmt.Println("Passed :) => IP:", ip, "/Key:", key, "/Wanted:", hn)
				judge = true
			}
		}
		if !judge {
			fmt.Println("Failed :( => IP:", ip, "/Key:", key, "/Wanted:", hn)
			t.Errorf("Error") // テストNGがあればエラーにする（エラー内容がわかるようにメッセージをちゃんと書く）
		}
	}
}
