package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URL         string
	Method      string
	function    func(w http.ResponseWriter, r *http.Request)
	RequireAuth bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URL, route.function).Methods(route.Method)
	}

	return r
}
