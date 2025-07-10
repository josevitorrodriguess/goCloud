package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (api *Api) getCallBackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	ctx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(ctx)

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("Erro na autenticação OAuth: %v", err)
		http.Error(w, "Erro na autenticação: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(w, "User:", user)

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}

func (api *Api) logoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")


	ctx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(ctx)

	err := gothic.Logout(w, r)
	if err != nil {
		log.Printf("Erro no logout: %v", err)
		http.Error(w, "Erro no logout", http.StatusInternalServerError)
		return
	}


	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}


func (api *Api) authHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")


	ctx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(ctx)


	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {

		log.Printf("Usuário já autenticado: %+v", gothUser)

		http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
		return
	}


	gothic.BeginAuthHandler(w, r)
}
