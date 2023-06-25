package router

import (
	"api/src/router/route"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return route.Configure(r)
}
