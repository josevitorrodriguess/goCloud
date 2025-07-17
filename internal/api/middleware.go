package api

import (
	"net/http"

	"github.com/markbates/goth/gothic"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := gothic.GetFromSession("user", r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Usuário não autenticado"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
