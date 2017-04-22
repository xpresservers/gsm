package testing

import (
	"testing"
)

func Assert(t *testing.T, expected interface{}, actual interface{}) bool {
	if expected != actual {
		Failed(t, expected, actual)
	}

	return true
}

func Failed(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf("Expected: '%s'; Actual: %s", expected, actual)
}

func Log(t *testing.T, message string) {
	t.Error(message)
}
