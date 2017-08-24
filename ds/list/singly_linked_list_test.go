package list_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/list"
)

func TestNewSinglyLinkedList(t *testing.T) {
	newList := list.NewSinglyLinkedList(4, 5, 6)
	if newList.Len() != 3 {
		t.Errorf("NewSinglyLinkedList: expected Len to be 3, got %d", newList.Len())
	}
	if !newList.Contains(4, 5, 6) {
		t.Error("NewSinglyLinkedList: expected Contains to be true, got false")
	}

	testInterfaceHelper(list.NewSinglyLinkedList, t)
}
