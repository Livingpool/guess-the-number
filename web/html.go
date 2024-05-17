package web

import (
	"html/template"
	"io"
)

func Render(w io.Writer, filepath string, data any) error {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, data); err != nil {
		return err
	}

	return nil
}
