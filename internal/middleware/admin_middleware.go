package middleware

import (
	"net/http"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value("role")
		if role == nil || role.(string) != "admin" {
			http.Error(w, "forbidden: admin only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
