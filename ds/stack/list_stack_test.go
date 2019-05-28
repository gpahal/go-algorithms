package stack_test

import (
	"testing"

	"github.com/gpahal/golib/ds/stack"
)

func TestListStack(t *testing.T) {
	testInterfaceHelper(t, stack.NewListStack)
}
