package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

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

const templateGlob = "/home/kimura/go-malproxy/server/templates/*.html"
const debug = true

var executor TemplateExecutor

func Server() {
	if debug {
		executor = DebugTemplateExecutor{templateGlob}

	} else {
		executor = ReleaseTemplateExecutor{
			template.Must(template.ParseGlob(templateGlob)),
		}
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/google", GoogleHandler)
	http.HandleFunc("/template", TemplateHandler)

	http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3000/
	executor.ExecuteTemplate(w, "index", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3000/hello
	fmt.Println("r->", r)
	executor.ExecuteTemplate(w, "hello", nil)
}

func GoogleHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3000/google
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	fmt.Println("検索ワード:", r.FormValue("params"))
	word := r.FormValue("params")
	res, err := service.GoogleSearch(word)
	if err != nil {
		executor.ExecuteTemplate(w, "err", nil)
	}
	fmt.Println("res:", res)
	err = service.ReadDataAndRewiteURL(res)
	if err != nil {
		executor.ExecuteTemplate(w, "err", nil)
	}
	file := "rewrite_" + res
	fmt.Println("file:", file)
	executor.ExecuteTemplate(w, file, nil)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3000/template?url=http://mitm.es3/amazon.co.jp
	fmt.Println("url", r.FormValue("url")) //取得したパラメータの表示
	url := r.FormValue("url")
	res, err := service.MainOperation(url)
	if err != nil {
		log.Fatal(err)
	}
	executor.ExecuteTemplate(w, res, nil)
}
