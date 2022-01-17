package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()
	e.Renderer = t
	e.Use(middleware.CORS())
	InitRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}

//http://localhost:1323/
