package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavomello-21/devbook/api/src/middleware"
)

type Routes struct {
	URL         string
	Method      string
	function    func(w http.ResponseWriter, r *http.Request)
	RequireAuth bool
}

func Config(r *mux.Router) *mux.Router {

	routes := userRoutes

	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URL, middleware.Logger(middleware.Authenticate(route.function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URL, middleware.Logger(route.function)).Methods(route.Method)
		}
	}

	r.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servidor funcionando"))
	}).Methods(http.MethodGet)

	return r
}
