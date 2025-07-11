package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/goCloud/internal/api"
)

func main() {
	godotenv.Load()
	api := api.Api{}
	api.Init()

	if err := http.ListenAndServe(":3000", api.Router); err != nil {
		panic(err)
	}
}
