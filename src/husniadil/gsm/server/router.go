package server

import (
	. "husniadil/gsm/handler"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/brands", GetAllBrands)
	router.HandleFunc("/devices", GetDeviceList).Queries("slug", "{slug:[a-z0-9\\-_]+}")

	return router
}
