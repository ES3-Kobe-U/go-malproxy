package handler

import (
	"fmt"
	"net/http"

	"github.com/go-malproxy/server/service"
)

func GoogleHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/google
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	fmt.Println("検索ワード:", r.FormValue("params"))
	word := r.FormValue("params")
	res, err := service.GoogleSearch(word)
	if err != nil {
		Executor.ExecuteTemplate(w, "err", nil)
	}
	fmt.Println("res:", res)
	err = service.RewriteUrlOfGoogleSearch(res)
	if err != nil {
		Executor.ExecuteTemplate(w, "err", nil)
	}
	file := "autogen_rewrite_" + res
	fmt.Println("file:", file)
	Executor.ExecuteTemplate(w, file, nil)
}
