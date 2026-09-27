package view

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path"
)

//go:embed templates
var files embed.FS

type Page string

type Fragment string

const (
	PageIndex    Page = "index.html"
	PageLogin    Page = "login.html"
	PageRegister Page = "register.html"
	PageLogout   Page = "logout.html"
	PageSessions Page = "sessions.html"
)

const (
	FragmentMe             Fragment = "me.html"
	FragmentSessionsMe     Fragment = "sessions_me.html"
	FragmentLoginErrors    Fragment = "login_errors.html"
	FragmentRegisterErrors Fragment = "register_errors.html"
)

var pageNames = []Page{PageIndex, PageLogin, PageRegister, PageLogout, PageSessions}

type View struct {
	pages     map[Page]*template.Template
	fragments *template.Template
}

type PageData struct {
	Authed  bool
	Current string
	Content any
}

func New() (*View, error) {
	fragments, err := template.ParseFS(files, "templates/fragments/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse fragments: %w", err)
	}

	paths, err := fs.Glob(files, "templates/partials/*.html")
	if err != nil {
		return nil, fmt.Errorf("glob pages: %w", err)
	}
	pages := make(map[Page]*template.Template, len(paths))
	for _, p := range paths {
		page := Page(path.Base(p))
		tmpl, err := template.ParseFS(files, "templates/layout.html", p)
		if err != nil {
			return nil, fmt.Errorf("parse page %s: %w", page, err)
		}
		pages[page] = tmpl
	}
	for _, page := range pageNames {
		if _, ok := pages[page]; !ok {
			return nil, fmt.Errorf("missing page template %s", page)
		}
	}

	return &View{pages: pages, fragments: fragments}, nil
}

func (v *View) Render(w io.Writer, f Fragment, data any) error {
	if err := v.fragments.ExecuteTemplate(w, string(f), data); err != nil {
		return fmt.Errorf("render fragment %s: %w", f, err)
	}
	return nil
}

func (v *View) RenderPage(w io.Writer, page Page, data PageData) error {
	tmpl, ok := v.pages[page]
	if !ok {
		return fmt.Errorf("unknown page %s", page)
	}
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		return fmt.Errorf("render page %s: %w", page, err)
	}
	return nil
}
