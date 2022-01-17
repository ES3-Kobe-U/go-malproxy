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
	templates: template.Must(template.ParseGlob("/home/kimura/go-malproxy/server/server/views/*.html")),
}

func InitRouter(e *echo.Echo) {
	e.GET("/", handler.AssignMainHandler)
	e.GET("/login", handler.AssignPostLoginData)
	e.GET("google-search/", handler.AssignGoogleSearchHandler)
}
