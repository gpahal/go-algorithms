package list_test

import (
	"testing"

	"github.com/gpahal/golib/ds/list"
)

func TestSinglyLinkedList(t *testing.T) {
	testInterfaceHelper(t, list.NewSinglyLinkedList)
}
