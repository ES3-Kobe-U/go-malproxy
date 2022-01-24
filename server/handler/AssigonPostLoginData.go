package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AssignPostLoginData(c echo.Context) error {
	params := new(LoginData)
	if err := c.Bind(params); err != nil {
		return err
	}

	log.Println("Email:", params.Email)
	log.Println("Password:", params.Password)
	return c.JSON(http.StatusOK, params)
}

// http://localhost:1323/login?Email=mail@example.com&Password=password