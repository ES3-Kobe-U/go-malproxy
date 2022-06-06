package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-malproxy/server/templates"
	"github.com/go-malproxy/server/usecase"
)

var ctx context.Context
var usecases usecase.Usecase = &usecase.Contents{IsAmazon: false, IsRakuten: false}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.Executor.ExecuteTemplate(w, "index", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	templates.Executor.ExecuteTemplate(w, "hello", nil)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\x1b[31mURL:\x1b[0m", r.FormValue("url"))
	url := r.FormValue("url")
	if strings.Contains(url, "https://www.amazon.co.jp/ap/signin?openid.pape.max_auth_age") {
		fmt.Println("Amazon Login")
		templates.Executor.ExecuteTemplate(w, "amazon-login", nil)
	} else {
		fmt.Println("Main Operation")
		res, err := usecase.MainOperation(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("res:", res)
		templates.Executor.ExecuteTemplate(w, res, nil)
	}
}
