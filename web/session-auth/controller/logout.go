package controller

import (
	"net/http"

	"session-auth/view"
)

func (c *Router) RenderLogoutConfirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.view.RenderPage(w, "logout.html", view.PageData{Authed: true, Current: "logout"}); err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
	}
}

func (c *Router) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie(sessionCookie); err == nil {
		if err := c.u.Logout(r.Context(), cookie.Value); err != nil {
			http.Error(w, "logout: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	clearSessionCookie(w)

	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusOK)
}

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
}
