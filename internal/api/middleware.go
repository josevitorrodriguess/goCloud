package api

import (
	"context"
	"net/http"

	"github.com/josevitorrodriguess/goCloud/internal/session"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, ok := session.GetSession(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Usuário não autenticado"))
			return
		}
		ctx := context.WithValue(r.Context(), "user_email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
