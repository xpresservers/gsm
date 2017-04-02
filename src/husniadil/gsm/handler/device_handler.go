package handler

import (
	"errors"
	"husniadil/gsm/service/device"
	. "husniadil/gsm/util/response"
	"net/http"
	"strconv"
)

func GetDeviceList(w http.ResponseWriter, req *http.Request) {
	slug := req.URL.Query().Get("slug")
	strPage := req.URL.Query().Get("page")

	page := 1
	if len(strPage) > 0 {
		p, err := strconv.Atoi(strPage)

		if err != nil || page < 1 {
			ResponseFailed(w, http.StatusBadRequest, errors.New("Page must be integer greater than zero"))
			return
		}

		page = p
	}

	devices, err := device.GetDeviceList(slug, page)

	if devices.CurrentPage > devices.TotalPage || devices.TotalPage == 0 {
		ResponseFailed(w, http.StatusNotFound, errors.New("Not found"))
		return
	}

	if err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err)
		return
	}

	ResponseOk(w, http.StatusOK, devices)
}
