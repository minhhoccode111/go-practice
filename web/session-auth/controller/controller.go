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
	mux.Handle("GET /", c.RequireAuth(http.HandlerFunc(c.Index)))
	mux.Handle("GET /sessions", c.RequireAuth(http.HandlerFunc(c.RenderSessions)))
	mux.Handle("GET /htmx/sessions/me", c.RequireAuth(http.HandlerFunc(c.HandleGetMySessions)))
	mux.Handle("DELETE /htmx/sessions/{id}", c.RequireAuth(http.HandlerFunc(c.HandleDeleteSession)))
	mux.Handle("DELETE /htmx/sessions", c.RequireAuth(http.HandlerFunc(c.HandleDeleteAllSessions)))
	mux.Handle("GET /htmx/me", c.RequireAuth(http.HandlerFunc(c.HandleGetMe)))
	mux.Handle("GET /logout", c.RequireAuth(http.HandlerFunc(c.RenderLogoutConfirm)))
	mux.Handle("GET /login", c.RedirectIfAuthed(http.HandlerFunc(c.RenderLogin)))
	mux.Handle("GET /register", c.RedirectIfAuthed(http.HandlerFunc(c.RenderRegister)))
	mux.HandleFunc("POST /htmx/login", c.HandleLogin)
	mux.HandleFunc("POST /htmx/register", c.HandleRegister)
	mux.HandleFunc("POST /htmx/logout", c.HandleLogout)
}

func (c *Router) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "index.html", view.PageData{Authed: true, Current: "home"}); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}
