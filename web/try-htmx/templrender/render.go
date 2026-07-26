package templrender

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(ctx context.Context, w http.ResponseWriter, component templ.Component, status ...int) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if len(status) > 0 {
		w.WriteHeader(status[0])
	}
	return component.Render(ctx, w)
}
