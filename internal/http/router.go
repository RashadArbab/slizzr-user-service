package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) newRouter() http.Handler {
	router := mux.NewRouter()
	router.Use(routeLog)
	router.Use(setHeaderJson)
	router.Use(setHeaderCors)

	router.HandleFunc("/users/single/{id}", getUser(server.Mongo)).Methods("GET")
	router.HandleFunc("/users/multiple", getMultipleUser(server.Mongo)).Methods("GET")
	return router
}
