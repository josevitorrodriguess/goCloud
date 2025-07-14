package api

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	api.Router.Get("/auth/{provider}/callback", api.getCallBackFunction)
	api.Router.Get("/auth/logout/{provider}", api.logoutHandler)
	api.Router.Get("/auth/{provider}", api.authHandler)
	api.Router.Put("/user/avatar", api.updateAvatarHandler)
	api.Router.Delete("/user", api.deleteUserHandler)
}
