package stack_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/stack"
)

func TestListStack(t *testing.T) {
	testInterfaceHelper(t, stack.New)
}
