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
		requestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		function:    controllers.FindById,
		requestAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodPut,
		function:    controllers.Update,
		requestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		function:    controllers.Delete,
		requestAuth: false,
	},
}
