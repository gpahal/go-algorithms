package set_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/set"
)

func TestNewNativeSet(t *testing.T) {
	testInterfaceHelper(t, set.NewNativeSet)
}
