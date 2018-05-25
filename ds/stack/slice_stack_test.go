package stack_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/stack"
)

func TestSliceStack(t *testing.T) {
	testInterfaceHelper(t, stack.NewSliceStack)
}
