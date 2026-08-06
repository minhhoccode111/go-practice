package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mhc/auth-go-react-zitadel/backend/internal"
)

func main() {
	issuer := envOr("ZITADEL_ISSUER", "http://localhost:8082")
	projectID := envOr("ZITADEL_PROJECT_ID", "")
	port := envOr("PORT", "8083")

	ctx := context.Background()
	verifier, err := internal.NewVerifier(ctx, issuer, projectID)
	if err != nil {
		slog.Error("cannot init verifier", "error", err)
		os.Exit(1)
	}
	slog.Info("verifier ready", "issuer", issuer)

	mux := http.NewServeMux()
	api := verifier.Middleware(http.HandlerFunc(handleMe))
	mux.Handle("/api/me", api)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           cors(mux),
		ReadHeaderTimeout: 10 * time.Second,
	}
	slog.Info("listening", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("server terminated", "error", err)
		os.Exit(1)
	}
}

func handleMe(w http.ResponseWriter, r *http.Request) {
	c, ok := internal.ClaimsFrom(r.Context())
	if !ok {
		internal.WriteErr(w, http.StatusUnauthorized, "unauthenticated")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(c)
}

// cors wraps a handler and adds dev CORS headers. Preflight requests
// for Authorization header are answered directly.
func cors(next http.Handler) http.Handler {
	allow := os.Getenv("CORS_ORIGIN")
	if allow == "" {
		allow = "http://localhost:5174"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if strings.HasPrefix(origin, "http://localhost:") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
