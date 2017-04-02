package server

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
