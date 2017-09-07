package stack

import "github.com/gpahal/go-algorithms/ds/list"

// Instance represents a stack instance implemented as a singly linked list.
type Instance struct {
	l *list.SinglyLinkedList
}

// New returns a new stack instance with the given items pushed into it.
func New(items ...int) *Instance {
	i := &Instance{l: &list.SinglyLinkedList{}}
	i.Push(items...)
	return i
}

// Length returns the number of items in the stack.
func (i *Instance) Len() int {
	return i.l.Len()
}

// Empty checks whether the stack is empty.
func (i *Instance) Empty() bool {
	return i.l.Empty()
}

// Clear deletes all the items from the stack.
func (i *Instance) Clear() {
	i.l.Clear()
}

// Values returns a slice of the items of the stack.
func (i *Instance) Values() []int {
	return i.l.Values()
}

// Push pushes the given items to the stack.
func (i *Instance) Push(items ...int) {
	i.l.PushFront(items...)
}

// Pop pops out an item from the stack in LIFO (Last In First Out) order.
func (i *Instance) Pop() (int, bool) {
	el := i.l.PopFront()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Copy creates a new copy of the stack.
func (i *Instance) Copy() *Instance {
	arr := i.Values()

	// reverse the values slice
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return New(arr...)
}
