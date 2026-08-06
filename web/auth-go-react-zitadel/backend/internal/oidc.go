package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
)

// Verifier validates access tokens signed by ZITADEL.
type Verifier struct {
	verifier *oidc.IDTokenVerifier
	aud      string // required project audience (urn:zitadel:...:project:id:<id>:aud)
}

// NewVerifier builds a token verifier from OIDC discovery + JWKS.
func NewVerifier(ctx context.Context, issuer string, projectID string) (*Verifier, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, fmt.Errorf("oidc discovery %s: %w", issuer, err)
	}

	// ZITADEL includes the raw project ID in the access token's aud claim,
	// so validate against it directly.
	config := &oidc.Config{
		SkipClientIDCheck: true,
		SupportedSigningAlgs: []string{
			"RS256", "RS384", "RS512",
			"PS256", "PS384", "PS512",
		},
	}

	return &Verifier{
		verifier: provider.Verifier(config),
		aud:      projectID,
	}, nil
}

// Claims is the user profile carried by a verified token.
type Claims struct {
	Subject   string            `json:"sub"`
	Email     string            `json:"email"`
	EmailVer  bool              `json:"email_verified"`
	Name      string            `json:"name"`
	GivenName string            `json:"given_name"`
	FamilyName string           `json:"family_name"`
	Username  string            `json:"preferred_username"`
	Aud       []string          `json:"aud"`
	Raw       map[string]any    `json:"-"`
}

// Verify parses + validates the Bearer token and returns its claims.
func (v *Verifier) Verify(ctx context.Context, token string) (*Claims, error) {
	idTok, err := v.verifier.Verify(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("verify token: %w", err)
	}

	var c Claims
	if err := idTok.Claims(&c); err != nil {
		return nil, fmt.Errorf("decode claims: %w", err)
	}
	// go-oidc enforces aud via ClientID when SkipClientIDCheck is false,
	// but assert explicitly anyway.
	if !hasAud(c.Aud, v.aud) {
		return nil, fmt.Errorf("token aud %v does not include %s", c.Aud, v.aud)
	}
	return &c, nil
}

func hasAud(auds []string, want string) bool {
	for _, a := range auds {
		if strings.EqualFold(a, want) {
			return true
		}
	}
	return false
}

// Middleware guards /api routes: requires a valid Bearer token and puts
// Claims into the request context.
type ctxKey int

const claimsKey ctxKey = 0

func ClaimsFrom(ctx context.Context) (*Claims, bool) {
	c, ok := ctx.Value(claimsKey).(*Claims)
	return c, ok
}

func (v *Verifier) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		token, ok := strings.CutPrefix(auth, "Bearer ")
		if !ok || token == "" {
			WriteErr(w, http.StatusUnauthorized, "missing bearer token")
			return
		}
		claims, err := v.Verify(r.Context(), token)
		if err != nil {
			slog.Warn("token verification failed", "error", err)
			WriteErr(w, http.StatusUnauthorized, "invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// WriteErr writes a JSON error response.
func WriteErr(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
