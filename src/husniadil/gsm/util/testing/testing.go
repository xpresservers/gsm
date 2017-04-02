package testing

import "testing"

func AssertFailed(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf("Expected: '%s'; Actual: %s", expected, actual)
}

func AssertFailedLog(t *testing.T, message string) {
	t.Error(message)
}
