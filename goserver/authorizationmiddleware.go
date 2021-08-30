package module

import (
	"log"
	"net/http"
	"strings"
)

type AuthorizationMiddleware struct {
	handler    http.Handler
	AuthToken  string
	allowLocal bool
}

func (aw *AuthorizationMiddleware) Handler(h http.Handler) http.Handler {
	aw.handler = h
	return aw
}

func (aw *AuthorizationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if aw.authorize(r) {
		aw.handler.ServeHTTP(w, r)
		return
	}
	log.Fatalf("Unauthorized request")
	http.Error(w, "Unauthorized request", http.StatusUnauthorized)
}

func (aw *AuthorizationMiddleware) authorize(r *http.Request) bool {
	if aw.allowLocal && strings.Index(r.RemoteAddr, "127.0.0.1") == 0 || strings.Index(r.RemoteAddr, "[::1]") == 0 {
		return true
	}
	authH := r.Header.Get("Authorization")
	token := strings.TrimLeft(authH, "Bearer: ")
	return token == aw.AuthToken
}
