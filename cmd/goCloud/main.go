package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/goCloud/internal/api"
)

func main() {
	godotenv.Load()
	api := api.Api{
		Router: chi.NewMux(),
	}

	if err := http.ListenAndServe(":3000", api.Router); err != nil {
		panic(err)
	}
}
