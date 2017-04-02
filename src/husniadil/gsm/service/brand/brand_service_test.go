package brand

import (
	"testing"

	assertion "husniadil/gsm/util/testing"
)

func TestGetAllBrands(t *testing.T) {
	brands, err := GetAllBrands()

	if err != nil {
		assertion.AssertFailedLog(t, "Failed to get brands")
	}

	firstBrandName := "Acer"
	if brands[0].Name != firstBrandName {
		assertion.AssertFailed(t, firstBrandName, brands[0].Name)
	}
}
