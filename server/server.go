package server

import (
	"net/http"
	"text/template"

	"github.com/go-malproxy/server/templates"
)

func Server() {
	if templates.Debug {
		templates.Executor = templates.DebugTemplateExecutor{Glob: templates.TemplateGlob}

	} else {
		templates.Executor = templates.ReleaseTemplateExecutor{
			Template: template.Must(template.ParseGlob(templates.TemplateGlob)),
		}
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/public/"))))
	InitRouter()
	http.ListenAndServe(":8081", nil)
}
