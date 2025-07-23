// package main

// import (
// 	"net/http"

// 	"github.com/joho/godotenv"
// 	"github.com/josevitorrodriguess/goCloud/internal/api"
// 	"github.com/josevitorrodriguess/goCloud/internal/logger"
// 	"github.com/josevitorrodriguess/goCloud/internal/repository"
// 	"github.com/josevitorrodriguess/goCloud/internal/storage/postgres"
// 	"github.com/josevitorrodriguess/goCloud/internal/usecase"
// )

// func main() {
// 	logger.Info("Starting goCloud server...")

// 	logger.Info("Loading environment variables")
// 	godotenv.Load()

// 	logger.Info("Connecting to database")
// 	db := postgres.ConnectDatabase()
// 	repo := repository.NewUserRepository(db)
// 	userUsecase := usecase.NewUserUsecase(*repo)

// 	logger.Info("Initializing API")
// 	apiInstance := api.Api{
// 		UserUsecase: userUsecase,
// 	}
// 	apiInstance.Init()

// 	logger.Info("Starting server on: http://localhost:3050")
// 	if err := http.ListenAndServe(":3050", apiInstance.Router); err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/goCloud/internal/api"
	"github.com/josevitorrodriguess/goCloud/internal/auth"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
	"github.com/josevitorrodriguess/goCloud/internal/repository"
	"github.com/josevitorrodriguess/goCloud/internal/session"
	"github.com/josevitorrodriguess/goCloud/internal/storage/aws"
	"github.com/josevitorrodriguess/goCloud/internal/storage/postgres"
	"github.com/josevitorrodriguess/goCloud/internal/usecase"
)

func main() {
	logger.Info("Starting goCloud server...")

	logger.Info("Loading environment variables")
	godotenv.Load()

	// Inicializa o Redis para sessão
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0 
	logger.Info("Conectando ao Redis em: %s", redisAddr)
	session.InitRedisSession(redisAddr, redisPassword, redisDB)

	logger.Info("Connecting to database")
	db := postgres.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(*repo)

	logger.Info("Inicializando S3 para FileUsecase...")
	s3Config := aws.GetS3ConfigFromEnv()
	s3Client, err := aws.NewS3Client(context.Background(), s3Config)
	if err != nil {
		logger.Error("Erro ao inicializar S3: %v", err)
		return
	}
	fileRepo := repository.NewFileRepository(s3Client, s3Config.Bucket)
	fileUsecase := usecase.NewFileUsecase(*fileRepo)

	logger.Info("Inicializando API...")
	apiInstance := api.Api{
		UserUsecase: userUsecase,
		FileUsecase: fileUsecase,
	}
	apiInstance.Init()

	auth.NewAuth()

	logger.Info("Servidor iniciado em: http://localhost:3050")
	if err := http.ListenAndServe(":3050", apiInstance.Router); err != nil {
		panic(err)
	}
}
