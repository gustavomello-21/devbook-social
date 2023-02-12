package router

import (
	"github.com/gorilla/mux"
	"github.com/gustavomello-21/devbook/api/src/router/routes"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router = routes.Config(router)

	return router
}
