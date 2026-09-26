package controller

import (
	"context"
	"errors"
	"net/http"

	"session-auth/entity"
)

type (
	userContextKey    struct{}
	sessionContextKey struct{}
)

func currentUser(r *http.Request) (*entity.User, bool) {
	u, ok := r.Context().Value(userContextKey{}).(*entity.User)
	return u, ok
}

func currentSession(r *http.Request) (*entity.Session, bool) {
	s, ok := r.Context().Value(sessionContextKey{}).(*entity.Session)
	return s, ok
}

func Sleep(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(time.Second)
		next.ServeHTTP(w, r)
	})
}

func (c *Router) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, s, err := c.authenticate(r)
		switch {
		case err != nil:
			http.Error(w, "auth: "+err.Error(), http.StatusInternalServerError)
		case u == nil:
			redirectToLogin(w, r)
		default:
			ctx := context.WithValue(r.Context(), userContextKey{}, u)
			ctx = context.WithValue(ctx, sessionContextKey{}, s)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func (c *Router) RedirectIfAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _, err := c.authenticate(r)
		switch {
		case err != nil:
			http.Error(w, "auth: "+err.Error(), http.StatusInternalServerError)
		case u != nil:
			http.Redirect(w, r, "/", http.StatusSeeOther)
		default:
			next.ServeHTTP(w, r)
		}
	})
}

func (c *Router) authenticate(r *http.Request) (*entity.User, *entity.Session, error) {
	cookie, err := r.Cookie(sessionCookie)
	if err != nil {
		return nil, nil, nil
	}
	u, s, err := c.u.Authenticate(r.Context(), cookie.Value)
	switch {
	case errors.Is(err, entity.ErrNotFound), errors.Is(err, entity.ErrInvalidCredentials):
		return nil, nil, nil
	case err != nil:
		return nil, nil, err
	}
	return u, s, nil
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", "/login")
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
