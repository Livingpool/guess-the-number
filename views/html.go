package views

import (
	"html/template"
	"io"
)

type TemplatesInterface interface {
	Render(w io.Writer, name string, data interface{}) error
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("./views/html/*.html")),
	}
}
