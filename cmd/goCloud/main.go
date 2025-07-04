package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/josevitorrodriguess/goCloud/internal/api"
)

func main() {
	api := api.Api{
		Router: chi.NewMux(),
	}

	if err := http.ListenAndServe(":8080", api.Router); err != nil {
		panic(err)
	}
}
