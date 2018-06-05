package list_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/list"
)

func TestDoublyLinkedList(t *testing.T) {
	testInterfaceHelper(t, list.NewDoublyLinkedList)
}
