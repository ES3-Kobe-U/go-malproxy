package server

import (
	"html/template"
	"net/http"

	"github.com/go-malproxy/server/handler"
)

func Server() {
	if handler.Debug {
		handler.Executor = handler.DebugTemplateExecutor{Glob: handler.TemplateGlob}

	} else {
		handler.Executor = handler.ReleaseTemplateExecutor{
			Template: template.Must(template.ParseGlob(handler.TemplateGlob)),
		}
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/public/"))))
	InitRouter()
	http.ListenAndServe(":3333", nil)
}
