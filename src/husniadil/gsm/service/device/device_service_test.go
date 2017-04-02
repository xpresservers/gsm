package device

import (
	"testing"

	assertion "husniadil/gsm/util/testing"
	"husniadil/gsm/service/brand"
)

func TestGetAllBrands(t *testing.T) {
	brands, err := brand.GetAllBrands()
	if err != nil {
		assertion.AssertFailedLog(t, "Failed to get brands")
	}

	acerTotalDevices := brands[0].NumberOfDevices
	perPage := 40
	totalPage := acerTotalDevices / perPage
	if acerTotalDevices % perPage > 0 {
		totalPage = totalPage + 1
	}

	devices, err := GetDeviceList("acer-phones-59", totalPage)

	if err != nil {
		assertion.AssertFailedLog(t, "Failed to get brands")
	}

	oldestDevice := devices.Items[len(devices.Items)-1]
	oldestDeviceSlug := "acer_dx900-2718"

	if oldestDeviceSlug != oldestDevice.Slug {
		assertion.AssertFailed(t, oldestDeviceSlug, oldestDevice.Slug)
	}
}
