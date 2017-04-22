package device

import (
	"testing"

	"husniadil/gsm/service/brand"
	assertion "husniadil/gsm/util/testing"
)

func TestGetDeviceList(t *testing.T) {
	// get all brands first to get a slug
	brands, err := brand.GetAllBrands()
	assertion.Assert(t, nil, err)

	acerTotalDevices := brands[0].NumberOfDevices
	perPage := 40
	totalPage := acerTotalDevices / perPage
	if acerTotalDevices%perPage > 0 {
		totalPage = totalPage + 1
	}

	// get device list
	devices, err := GetDeviceList("acer-phones-59", totalPage)
	assertion.Assert(t, nil, err)

	oldestDevice := devices.Items[len(devices.Items)-1]
	oldestDeviceSlug := "acer_dx900-2718"
	assertion.Assert(t, oldestDeviceSlug, oldestDevice.Slug)
}
