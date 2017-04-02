package handler

import (
	"husniadil/gsm/service/brand"
	. "husniadil/gsm/util/response"
	"net/http"
)

func GetAllBrands(w http.ResponseWriter, req *http.Request) {
	brands, err := brand.GetAllBrands()

	if err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err)
		return
	}

	ResponseOk(w, http.StatusOK, brands)
}
