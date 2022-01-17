package server

import (
	"github.com/go-malproxy/server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.CORS())
	//Router
	e.GET("/login", handler.AssignPostLoginData)
	e.GET("google-search/", handler.AssignGoogleSearchHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
