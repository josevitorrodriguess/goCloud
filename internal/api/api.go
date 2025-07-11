package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/josevitorrodriguess/goCloud/internal/auth"
)

type Api struct {
	Router *chi.Mux
}

func (api *Api) Init() {
	api.Router = chi.NewMux()
	auth.NewAuth()
	api.NewRouter()
}
