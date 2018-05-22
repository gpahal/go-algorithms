package queue

import (
	"github.com/gpahal/go-algorithms/ds/list"
)

// Queue represents a queue instance implemented as a doubly linked list.
type Queue struct {
	l *list.DoublyLinkedList
}

// New returns a new queue instance with the given items enqueued/added to it at the end.
func New(items ...int) *Queue {
	q := &Queue{l: &list.DoublyLinkedList{}}
	q.Enqueue(items...)
	return q
}

// Len returns the number of items in the queue.
func (q *Queue) Len() int {
	return q.l.Len()
}

// Empty checks whether the queue is empty.
func (q *Queue) Empty() bool {
	return q.l.Empty()
}

// Clear deletes all the items from the queue.
func (q *Queue) Clear() {
	q.l.Clear()
}

// Values returns a slice of the items of the queue.
func (q *Queue) Values() []int {
	return q.l.Values()
}

// Front returns the front/oldest enqueued element of the stack. If the queue is empty, second return value is false.
func (q *Queue) Front() (int, bool) {
	el := q.l.First()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Enqueue adds the items at the end of the queue.
func (q *Queue) Enqueue(items ...int) {
	q.l.PushBack(items...)
}

// Dequeue removes the item from the front of the queue and returns it. If the queue is empty, second return value is
// false.
func (q *Queue) Dequeue() (int, bool) {
	el := q.l.PopFront()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Copy creates a new copy of the queue.
func (q *Queue) Copy() *Queue {
	return New(q.Values()...)
}
