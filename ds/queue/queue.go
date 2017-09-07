package queue

import "github.com/gpahal/go-algorithms/ds/list"

// Instance represents a queue instance implemented as a doubly linked list.
type Instance struct {
	l *list.DoublyLinkedList
}

// New returns a new queue instance with the given items enqueued/added to it at the end.
func New(items ...int) *Instance {
	i := &Instance{l: &list.DoublyLinkedList{}}
	i.Enqueue(items...)
	return i
}

// Length returns the number of items in the queue.
func (i *Instance) Len() int {
	return i.l.Len()
}

// Empty checks whether the queue is empty.
func (i *Instance) Empty() bool {
	return i.l.Empty()
}

// Clear deletes all the items from the queue.
func (i *Instance) Clear() {
	i.l.Clear()
}

// Values returns a slice of the items of the queue.
func (i *Instance) Values() []int {
	return i.l.Values()
}

// Front returns the front/oldest enqueued element of the stack.
func (i *Instance) Front() (int, bool) {
	el := i.l.First()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Enqueue adds the items at the end of the queue.
func (i *Instance) Enqueue(items ...int) {
	i.l.PushBack(items...)
}

// Dequeue removes the item from the front of the queue and returns it.
func (i *Instance) Dequeue() (int, bool) {
	el := i.l.PopFront()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Copy creates a new copy of the queue.
func (i *Instance) Copy() *Instance {
	return New(i.Values()...)
}
