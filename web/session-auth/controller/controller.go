package controller

import (
	"net/http"

	"session-auth/config"
	"session-auth/service"
	"session-auth/view"
)

type Router struct {
	cfg  *config.Config
	u    service.User
	view *view.View
}

func NewRouter(cfg *config.Config, u service.User, v *view.View) *Router {
	return &Router{
		cfg:  cfg,
		u:    u,
		view: v,
	}
}

func (c *Router) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", c.Index)
	mux.HandleFunc("GET /login", c.RenderLogin)
	mux.HandleFunc("POST /login", c.HandleLogin)
	mux.HandleFunc("GET /register", c.RenderRegister)
	mux.HandleFunc("POST /register", c.HandleRegister)
}

func (c *Router) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "index.html", nil); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}
