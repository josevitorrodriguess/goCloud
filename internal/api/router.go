package api
import "github.com/go-chi/chi/v5/middleware"

func (api *Api) NewRouter() {
	api.Router.Use(middleware.Logger)

	api.Router.Get("/auth/{provider}/callback", api.getCallBackFunction)
	api.Router.Get("/auth/logout/{provider}", api.logoutHandler)
	api.Router.Get("/auth/{provider}", api.authHandler)

}
