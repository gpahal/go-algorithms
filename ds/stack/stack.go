package stack

import (
	"github.com/gpahal/go-algorithms/ds/list"
)

// Stack represents a stack instance implemented as a singly linked list.
type Stack struct {
	l *list.SinglyLinkedList
}

// New returns a new stack instance with the given items pushed into it.
func New(items ...int) *Stack {
	s := &Stack{l: &list.SinglyLinkedList{}}
	s.Push(items...)
	return s
}

// Len returns the number of items in the stack.
func (s *Stack) Len() int {
	return s.l.Len()
}

// Empty checks whether the stack is empty.
func (s *Stack) Empty() bool {
	return s.l.Empty()
}

// Clear deletes all the items from the stack.
func (s *Stack) Clear() {
	s.l.Clear()
}

// Values returns a slice of the items of the stack.
func (s *Stack) Values() []int {
	return s.l.Values()
}

// Top returns the top/last pushed element of the stack. If the stack is empty, second return value is false.
func (s *Stack) Top() (int, bool) {
	el := s.l.First()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Push pushes the given items to the stack.
func (s *Stack) Push(items ...int) {
	s.l.PushFront(items...)
}

// Pop pops out an item from the stack in LIFO (Last In First Out) order. If the stack is empty, second return value is
// false.
func (s *Stack) Pop() (int, bool) {
	el := s.l.PopFront()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Copy creates a new copy of the stack.
func (s *Stack) Copy() *Stack {
	arr := s.Values()

	// reverse the values slice
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return New(arr...)
}
