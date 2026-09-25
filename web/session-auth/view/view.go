package view

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed templates
var files embed.FS

type View struct {
	tmpl *template.Template
}

func New() (*View, error) {
	tmpl, err := template.ParseFS(files,
		"templates/*.html",
		"templates/partials/*.html",
		"templates/fragments/*.html",
	)
	if err != nil {
		return nil, fmt.Errorf("error new view: %w", err)
	}
	return &View{tmpl: tmpl}, nil
}

func (v *View) Render(w io.Writer, name string, data any) error {
	return v.tmpl.ExecuteTemplate(w, name, data)
}

func (v *View) RenderPage(w io.Writer, page string, data any) error {
	tmpl, err := template.ParseFS(files, "templates/layout.html", "templates/partials/"+page)
	if err != nil {
		return fmt.Errorf("parse page %s: %w", page, err)
	}
	return tmpl.ExecuteTemplate(w, "layout", data)
}
