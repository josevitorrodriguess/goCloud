package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/goCloud/internal/api"
	"github.com/josevitorrodriguess/goCloud/internal/repository"
	"github.com/josevitorrodriguess/goCloud/internal/storage/postgres"
	"github.com/josevitorrodriguess/goCloud/internal/usecase"
)

func main() {
	godotenv.Load()
	// Conecta ao banco e instancia o repositório e usecase
	db := postgres.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(*repo)

	apiInstance := api.Api{
		UserUsecase: userUsecase,
	}
	apiInstance.Init()

	if err := http.ListenAndServe(":3050", apiInstance.Router); err != nil {
		panic(err)
	}
}
