package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func HelloHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", nil)
}
