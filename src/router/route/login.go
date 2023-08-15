package route

import (
	"api/src/controllers"
	"net/http"
)

var loginRouter = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	function:    controllers.Login,
	requestAuth: false,
}
