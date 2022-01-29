package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/go-malproxy/server/handler"
	"github.com/go-malproxy/server/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var MainTemplate = &Template{
	templates: template.Must(template.ParseGlob("/home/kimura/go-malproxy/server/templates/*.html")),
}

func EchoServer() {
	e := echo.New()
	e.Renderer = MainTemplate
	fmt.Println(MainTemplate.templates.Name())
	e.Use(middleware.CORS())
	InitRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func InitRouter(e *echo.Echo) {
	e.GET("/", EchoIndexHandler)
	e.GET("/hello", EchoHelloHandler)

	e.GET("/template", EchoTemplateHandler)

	e.GET("/login", handler.PostLoginData)
	e.GET("/google-search", handler.GoogleSearchHandler)
}

func EchoIndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func EchoHelloHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", nil)
}

func EchoTemplateHandler(c echo.Context) error { // http://localhost:1323/template?url=https://amazon.co.jp
	var params struct {
		URL string `json:"url"`
	}
	if err := c.Bind(&params); err != nil {
		log.Fatal(err)
		return c.Render(http.StatusInternalServerError, "err", nil)
	}
	res, err := service.MainOperation(params.URL)
	if err != nil {
		log.Fatal(err)
		return c.Render(http.StatusInternalServerError, "err", nil)
	}
	return c.Render(http.StatusOK, res, nil)
}
