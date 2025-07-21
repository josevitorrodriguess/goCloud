package api

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/josevitorrodriguess/goCloud/internal/jsonutils"
)

func (api *Api) NewRouter() {
	api.Router.Use(middleware.Logger)

	// Middleware de CORS para permitir frontend em localhost:5173
	api.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	api.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		jsonutils.EncodeJson(w, r, http.StatusOK, "TUDO CERTO")
	})
	api.Router.Get("/auth/{provider}/callback", api.getCallBackFunction)
	api.Router.With(authMiddleware).Get("/auth/logout/{provider}", api.logoutHandler)
	api.Router.Get("/auth/{provider}", api.authHandler)
	// api.Router.Put("/user/avatar", api.updateAvatarHandler)
	// api.Router.Delete("/user", api.deleteUserHandler)

	// Rotas de arquivos
	api.Router.With(authMiddleware).Post("/file/upload", api.UploadFileHandler)
	api.Router.With(authMiddleware).Get("/file/download", api.DownloadFileHandler)
	api.Router.With(authMiddleware).Get("/file/list", api.ListFilesHandler)
	api.Router.With(authMiddleware).Delete("/file/delete", api.DeleteFileHandler)
}
