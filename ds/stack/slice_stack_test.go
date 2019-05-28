package stack_test

import (
	"testing"

	"github.com/gpahal/golib/ds/stack"
)

func TestSliceStack(t *testing.T) {
	testInterfaceHelper(t, stack.NewSliceStack)
}
