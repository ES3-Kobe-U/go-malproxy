package server

import (
	"net/http"

	"github.com/go-malproxy/server/handler"
)

func InitRouter() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/google", handler.GoogleHandler)
	http.HandleFunc("/rakuten-login", handler.RakutenLoginHandler)
	http.HandleFunc("/rakuten-login-info", handler.RakutenHandler)
	http.HandleFunc("/amazon-login", handler.AmazonLoginHandler)
	http.HandleFunc("/amazon-login-info", handler.AmazonHandler)
	http.HandleFunc("/template", handler.TemplateHandler)
}
