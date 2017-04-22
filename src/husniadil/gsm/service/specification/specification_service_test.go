package specification

import (
	"testing"

	assertion "husniadil/gsm/util/testing"
)

func TestGetDeviceList(t *testing.T) {
	specs, err := GetSpecification("acer_liquid_z4-6115")

	assertion.Assert(t, nil, err)
	assertion.Assert(t, "Acer", specs.Brand)
	assertion.Assert(t, "Acer Liquid Z4", specs.DeviceName)
	assertion.Assert(t, "Phone", specs.DeviceType)
}
