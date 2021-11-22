package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//server.Run() //TODO: 後々echoに移行したい
	write, err := os.Create("index.html") // カレントディレクトリにindex.htmlを作成
	if err != nil {
		log.Fatal(err)
	}

	read, err := os.Open("server/handler/output/index.html") // server/handler/output/index.htmlを読み込み
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(write, read) // コピー
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	log.Fatal(http.ListenAndServe(":1323", nil)) //起動
	// http://localhost:1323/static/index.html
}
