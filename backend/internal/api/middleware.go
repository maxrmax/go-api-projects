package api

import "net/http"

type authMiddleware struct {
	next http.Handler
}

func AuthMiddleware(next http.Handler) http.Handler {
	return authMiddleware{next: next}
}

func (m authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token != "Bearer demo-token" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	m.next.ServeHTTP(w, r)
}
