package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func AssignIndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
