package controller

import (
	"errors"
	"net/http"

	"session-auth/entity"
)

func (c *Router) RenderRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "register.html", nil); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}

func (c *Router) HandleRegister(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm := r.FormValue("repeat-password")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if password != confirm {
		c.registerErrors(w, "passwords do not match")
		return
	}

	_, err := c.u.Register(r.Context(), email, password)
	switch {
	case errors.Is(err, entity.ErrInvalidEmail):
		c.registerErrors(w, "enter a valid email address")
	case errors.Is(err, entity.ErrWeakPassword):
		c.registerErrors(w, "password must be at least 8 characters")
	case errors.Is(err, entity.ErrEmailTaken):
		c.registerErrors(w, "that email is already registered")
	case err != nil:
		http.Error(w, "register: "+err.Error(), http.StatusInternalServerError)
	default:
		if err := c.view.Render(w, "register_success.html", nil); err != nil {
			http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func (c *Router) registerErrors(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	if err := c.view.Render(w, "register_errors.html", msg); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}
