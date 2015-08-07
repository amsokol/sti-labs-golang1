package controllers

import (
	"html/template"
	"io"
)

type (
// Template provides HTML template rendering
	Renderer struct {
		templates *template.Template
	}
)

// Render HTML
func (t *Renderer) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func CreateRenderer() *Renderer {
	return &Renderer{
		// Cached templates
		templates: template.Must(template.ParseFiles(
			Page1File,
			Page2File)),
	}
}
