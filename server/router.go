package server

import (
	"io"
	"text/template"

	"github.com/go-malproxy/server/handler"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.Execute(w, nil)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("/home/kimura/go-malproxy/server/templates/*.html")),
}

func InitRouter(e *echo.Echo) {
	e.GET("/", handler.IndexHandler)
	e.GET("/login", handler.PostLoginData)
	e.GET("google-search/", handler.GoogleSearchHandler)
}

// http://localhost:1323/
