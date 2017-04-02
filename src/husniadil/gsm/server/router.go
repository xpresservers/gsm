package server

import (
	. "husniadil/gsm/handler"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/brands", GetAllBrands)

	return router
}
