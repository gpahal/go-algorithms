package list_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/list"
)

func TestSinglyLinkedList(t *testing.T) {
	testInterfaceHelper(t, list.NewSinglyLinkedList)
}
