package handler

import (
	"fmt"
	"net/http"

	"github.com/go-malproxy/server/templates"
	"github.com/go-malproxy/server/usecase"
)

func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	fmt.Println("params:", r.FormValue("params"))
	word := r.FormValue("params")
	res, err := usecase.GoogleSearch(word)
	if err != nil {

		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	fmt.Println("res:", res)
	err = usecase.RewriteUrlOfGoogleSearch(res)
	if err != nil {

		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	file := "autogen_rewrite_" + res
	templates.Executor.ExecuteTemplate(w, file, nil)
}
