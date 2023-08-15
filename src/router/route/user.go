package route

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/user",
		Method:      http.MethodPost,
		function:    controllers.Create,
		requestAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodGet,
		function:    controllers.FindByFilter,
		requestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		function:    controllers.FindById,
		requestAuth: true,
	},
	{
		URI:         "/user",
		Method:      http.MethodPut,
		function:    controllers.Update,
		requestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		function:    controllers.Delete,
		requestAuth: true,
	},
}
