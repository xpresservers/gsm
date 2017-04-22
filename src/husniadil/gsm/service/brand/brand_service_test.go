package brand

import (
	assertion "husniadil/gsm/util/testing"
	"testing"
)

func TestGetAllBrands(t *testing.T) {
	brands, err := GetAllBrands()
	assertion.Assert(t, nil, err)

	firstBrandName := "Acer"
	assertion.Assert(t, firstBrandName, brands[0].Name)
}
