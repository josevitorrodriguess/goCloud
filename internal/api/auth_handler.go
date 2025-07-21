package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/josevitorrodriguess/goCloud/internal/domain"
	"github.com/josevitorrodriguess/goCloud/internal/jsonutils"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
	"github.com/josevitorrodriguess/goCloud/internal/session"
	"github.com/markbates/goth/gothic"
)

type contextKey string

func (api *Api) getCallBackFunction(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("SCS: contexto de sessão NÃO encontrado no callback! (panic capturado)")
		}
	}()
	provider := chi.URLParam(r, "provider")

	ctx := context.WithValue(r.Context(), contextKey("provider"), provider)
	r = r.WithContext(ctx)

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		logger.Error("Erro na autenticação OAuth: %v", err)
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": "Authentication failed"})
		return
	}

	logger.Info("Usuário autenticado: %s (%s)", gothUser.Name, gothUser.Email)

	// Loga o conteúdo da sessão após autenticação
	sessUser, err := gothic.GetFromSession("user", r)
	if err != nil {
		logger.Error("[DEBUG] Não foi possível recuperar 'user' da sessão após login: %v", err)
	} else {
		logger.Info("[DEBUG] Conteúdo da sessão 'user' após login: %s", sessUser)
	}

	// Cria sessão no Redis
	_, err = session.SetSession(w, gothUser.Email)
	if err != nil {
		logger.Error("Erro ao criar sessão Redis: %v", err)
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": "Erro ao criar sessão"})
		return
	}

	// Salva/atualiza usuário
	api.UserUsecase.SaveUser(&domain.User{
		Provider:   provider,
		ProviderID: gothUser.UserID,
		Name:       gothUser.Name,
		Email:      gothUser.Email,
		AvatarURL:  gothUser.AvatarURL,
	})

	http.Redirect(w, r, "http://localhost:3050", http.StatusFound)
}

func (api *Api) logoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	ctx := context.WithValue(r.Context(), contextKey("provider"), provider)
	r = r.WithContext(ctx)

	err := gothic.Logout(w, r)
	if err != nil {
		logger.Error("Erro no logout: %v", err)
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": "Logout failed"})
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (api *Api) authHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Entrou no authHandler para iniciar autenticação")
	provider := chi.URLParam(r, "provider")

	ctx := context.WithValue(r.Context(), contextKey("provider"), provider)
	r = r.WithContext(ctx)

	gothic.BeginAuthHandler(w, r)
}
