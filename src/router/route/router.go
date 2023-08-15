package route

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	function    func(http.ResponseWriter, *http.Request)
	requestAuth bool
}

// Add all routes into router
func Configure(r *mux.Router) *mux.Router {
	routers := userRoutes
	routers = append(routers, loginRouter)
	for _, route := range routers {
		if route.requestAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Autenticar(route.function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.function)).Methods(route.Method)
		}
	}
	return r
}
