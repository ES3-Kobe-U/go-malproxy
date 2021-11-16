package server

import (
	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()
	Router(e) //router
	e.Logger.Fatal(e.Start(":1323"))
}
