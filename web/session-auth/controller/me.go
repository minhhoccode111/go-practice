package controller

import (
	"net/http"

	"session-auth/entity"
)

func (c *Router) HandleGetMe(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(userContextKey{}).(*entity.User)
	if !ok {
		http.Error(w, "me: no user in context", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.Render(w, "me.html", u.Email); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}
