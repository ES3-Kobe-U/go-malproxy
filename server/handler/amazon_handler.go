package handler

import (
	"fmt"
	"log"
	"net/http"
)

func AmazonLoginHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/amazon-login
	Executor.ExecuteTemplate(w, "amazon-login", nil)
}

func AmazonHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/amazon-login-info
	fmt.Println("method:", r.Method)
	fmt.Println("email", r.FormValue("email"))
	fmt.Println("password", r.FormValue("password"))
	email := r.FormValue("email")
	password := r.FormValue("password")
	err := services.CheckingTheIntegrityOfAmazonInformation(email, password)
	if err != nil {
		Executor.ExecuteTemplate(w, "err", nil)
	}
	if err := services.CheckingContextContents(); err != nil {
		log.Fatal(err)
	}
	Executor.ExecuteTemplate(w, "autogen_amazon_info", nil)
}
