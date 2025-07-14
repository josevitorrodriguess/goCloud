package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/josevitorrodriguess/goCloud/internal/auth"
	"github.com/josevitorrodriguess/goCloud/internal/usecase"
)

type Api struct {
	Router      *chi.Mux
	UserUsecase *usecase.UserUsecase
}

func (api *Api) Init() {
	api.Router = chi.NewMux()
	auth.NewAuth()
	api.NewRouter()
}
