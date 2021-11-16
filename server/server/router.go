package server

import (
	handler "github.com/go-malproxy/server/handler/linux"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.GET("/auth", handler.AssignLinuxAuthHander)
}
