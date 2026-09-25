package controller

import "net/http"

func (c *Router) RenderLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "login.html", nil); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}

func (c *Router) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "index.html", nil); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}
