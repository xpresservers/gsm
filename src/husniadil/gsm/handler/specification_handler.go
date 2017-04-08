package handler

import (
	"husniadil/gsm/service/specification"
	. "husniadil/gsm/util/response"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSpecification(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	slug := vars["slug"]

	specs, err := specification.GetSpecification(slug)

	if err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err)
		return
	}

	ResponseOk(w, http.StatusOK, specs)
}
