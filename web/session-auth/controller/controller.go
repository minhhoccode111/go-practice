package controller

import (
	"net/http"

	"session-auth/config"
	"session-auth/service"
)

type Router struct {
	cfg *config.Config
	u   service.User
}

func NewRouter(cfg *config.Config, u service.User) *Router {
	return &Router{
		cfg: cfg,
		u:   u,
	}
}

func (c *Router) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", c.Index)
}

func (c *Router) Index(w http.ResponseWriter, r *http.Request) {
}
