package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/goCloud/internal/api"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
	"github.com/josevitorrodriguess/goCloud/internal/repository"
	"github.com/josevitorrodriguess/goCloud/internal/storage/postgres"
	"github.com/josevitorrodriguess/goCloud/internal/usecase"
)

func main() {
	logger.Info("Starting goCloud server...")
	
	logger.Info("Loading environment variables")
	godotenv.Load()
	
	logger.Info("Connecting to database")
	db := postgres.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(*repo)

	logger.Info("Initializing API")
	apiInstance := api.Api{
		UserUsecase: userUsecase,
	}
	apiInstance.Init()

	logger.Info("Starting server on: http://localhost:3050")
	if err := http.ListenAndServe(":3050", apiInstance.Router); err != nil {
		panic(err)
	}
}
