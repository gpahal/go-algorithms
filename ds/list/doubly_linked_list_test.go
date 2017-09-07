package list_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/list"
)

func TestNewDoublyLinkedList(t *testing.T) {
	newList := list.NewDoublyLinkedList(4, 5, 6)
	if newList.Len() != 3 {
		t.Errorf("NewDoublyLinkedList: expected Len to be 3, got %d", newList.Len())
	}
	if !newList.Contains(4, 5, 6) {
		t.Error("NewDoublyLinkedList: expected Contains to be true, got false")
	}

	testInterfaceHelper(list.NewDoublyLinkedList, t)
}
