package route

import (
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

	for _, route := range routers {
		r.HandleFunc(route.URI, route.function).Methods(route.Method)
	}
	return r
}
