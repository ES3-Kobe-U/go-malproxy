package server

import (
	"github.com/go-malproxy/server/handler"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.GET("/login", handler.AssignPostLoginData)
	e.GET("/google-search", handler.AssignGoogleSearchHandler)
}
