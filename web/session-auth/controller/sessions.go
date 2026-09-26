package controller

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"session-auth/entity"
	"session-auth/view"
)

type sessionView struct {
	ID        uint
	ExpiresAt time.Time
	Current   bool
}

func (c *Router) RenderSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "sessions.html", view.PageData{Authed: true, Current: "sessions"}); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}

func (c *Router) HandleGetMySessions(w http.ResponseWriter, r *http.Request) {
	u, ok := currentUser(r)
	if !ok {
		http.Error(w, "sessions: no user in context", http.StatusInternalServerError)
		return
	}

	sessions, err := c.u.ListSessions(r.Context(), u.ID)
	if err != nil {
		http.Error(w, "sessions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	current, _ := currentSession(r)
	rows := make([]sessionView, 0, len(sessions))
	for _, s := range sessions {
		rows = append(rows, sessionView{
			ID:        s.ID,
			ExpiresAt: s.ExpiresAt,
			Current:   current != nil && s.ID == current.ID,
		})
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.Render(w, "sessions_me.html", rows); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}

func (c *Router) HandleDeleteSession(w http.ResponseWriter, r *http.Request) {
	u, ok := currentUser(r)
	if !ok {
		http.Error(w, "sessions: no user in context", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "sessions: invalid session id", http.StatusBadRequest)
		return
	}

	if current, ok := currentSession(r); ok && current.ID == uint(id) {
		http.Error(w, "sessions: cannot delete the current session, use /logout", http.StatusForbidden)
		return
	}

	err = c.u.DeleteSession(r.Context(), u.ID, uint(id))
	if err != nil && !errors.Is(err, entity.ErrNotFound) {
		http.Error(w, "sessions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Router) HandleDeleteAllSessions(w http.ResponseWriter, r *http.Request) {
	u, ok := currentUser(r)
	if !ok {
		http.Error(w, "sessions: no user in context", http.StatusInternalServerError)
		return
	}

	if err := c.u.DeleteAllSessions(r.Context(), u.ID); err != nil {
		http.Error(w, "sessions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	clearSessionCookie(w)
	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusOK)
}
