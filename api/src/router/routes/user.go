package routes

import (
	"net/http"

	"github.com/gustavomello-21/devbook/api/src/controllers"
)

var userRoutes = []Routes{
	{
		URL:         "/usuarios",
		Method:      http.MethodPost,
		function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URL:         "/usuarios",
		Method:      http.MethodGet,
		function:    controllers.GetAllUsers,
		RequireAuth: false,
	},
	{
		URL:         "/usuarios/{userId}",
		Method:      http.MethodGet,
		function:    controllers.GetUserById,
		RequireAuth: false,
	},
	{
		URL:         "/usuarios/{userId}",
		Method:      http.MethodPut,
		function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URL:         "/usuarios/{userId}",
		Method:      http.MethodDelete,
		function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
