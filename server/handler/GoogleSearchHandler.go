package handler

import (
	"log"
	"net/http"

	service "github.com/go-malproxy/server/service"
	"github.com/labstack/echo"
)

type GoogleSearchParams struct {
	Params string `json:"params"`
}

func GoogleSearchHandler(c echo.Context) error {
	Query := new(GoogleSearchParams)
	if err := c.Bind(Query); err != nil {
		return err
	}
	log.Println(Query.Params)
	err := service.GoogleSearch(Query.Params)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusOK, nil) //TODO:ここ変更する必要あり
}

//http://localhost:1323/google-search?params=神戸 大学
