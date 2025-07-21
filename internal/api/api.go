package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/josevitorrodriguess/goCloud/internal/auth"
	"github.com/josevitorrodriguess/goCloud/internal/usecase"
)

type Api struct {
	Router      *chi.Mux
	UserUsecase *usecase.UserUsecase
	FileUsecase *usecase.FileUsecase
}

func (api *Api) Init() {
	api.Router = chi.NewMux()
	auth.NewAuth()
	// api.Router.Use(auth.Manager.LoadAndSave) // Removido: sessão agora é só Redis
	api.NewRouter()
}
