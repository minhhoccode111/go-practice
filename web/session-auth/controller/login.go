package controller

import (
	"errors"
	"net/http"

	"session-auth/entity"
	"session-auth/view"
)

const sessionCookie = "session"

func (c *Router) RenderLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct{ Registered bool }{
		Registered: r.URL.Query().Get("registered") == "1",
	}
	if err := c.view.RenderPage(w, "login.html", view.PageData{Current: "login", Content: data}); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}

func (c *Router) HandleLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		c.loginErrors(w, "enter your email and password")
		return
	}

	token, err := c.u.Login(r.Context(), email, password)
	switch {
	case errors.Is(err, entity.ErrNotFound), errors.Is(err, entity.ErrInvalidCredentials):
		c.loginErrors(w, "invalid email or password")
	case err != nil:
		http.Error(w, "login: "+err.Error(), http.StatusInternalServerError)
	case token == "":
		http.Error(w, "login: empty session token", http.StatusInternalServerError)
	default:
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookie,
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
	}
}

func (c *Router) loginErrors(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	if err := c.view.Render(w, "login_errors.html", msg); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}
