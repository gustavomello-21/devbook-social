package routes

import (
	"net/http"

	"github.com/gustavomello-21/devbook/api/src/controllers"
)

var loginRoute = Routes{
	URL:         "/login",
	Method:      http.MethodPost,
	function:    controllers.Login,
	RequireAuth: false,
}
