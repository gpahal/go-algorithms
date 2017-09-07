package queue_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/queue"
)

func TestNew(t *testing.T) {
	newQueue := queue.New(4, 5, 6)
	if newQueue.Len() != 3 {
		t.Errorf("New 4, 5, 6: expected Len to be 3, got %d", newQueue.Len())
	}

	tmpArr := newQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6}) {
		t.Errorf("New 4, 5, 6: expected Values to be [4 5 6], got %v", tmpArr)
	}
}

func TestInstance_Len(t *testing.T) {
	newQueue := queue.New()
	if newQueue.Len() != 0 {
		t.Errorf("Len: expected Len to be 0, got %d", newQueue.Len())
	}

	newQueue.Enqueue(4, 5, 6)
	if newQueue.Len() != 3 {
		t.Errorf("Len: expected Len to be 3, got %d", newQueue.Len())
	}
}

func TestInstance_Empty(t *testing.T) {
	newQueue := queue.New()
	if !newQueue.Empty() {
		t.Error("Empty: expected Empty to be true, got false")
	}

	newQueue.Enqueue(4, 5, 6)
	if newQueue.Empty() {
		t.Error("Empty: expected Empty to be false, got true")
	}
}

func TestInstance_Clear(t *testing.T) {
	newQueue := queue.New(4, 5, 6)
	newQueue.Clear()
	if newQueue.Len() != 0 {
		t.Errorf("Clear: expected Len to be 0, got %d", newQueue.Len())
	}
}

func TestInstance_Values(t *testing.T) {
	newQueue := queue.New()
	tmpArr := newQueue.Values()
	if len(tmpArr) != 0 {
		t.Errorf("Values: expected Values to be [], got %#v", tmpArr)
	}

	newQueue.Enqueue(4, 5, 6)
	tmpArr = newQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6}) {
		t.Errorf("Values: expected Values to be [4 5 6], got %v", tmpArr)
	}
}

func TestInstance_Front(t *testing.T) {
	newQueue := queue.New(4, 5, 6)
	newQueue.Enqueue(7)
	val, ok := newQueue.Front()
	if !ok || val != 4 {
		t.Errorf("Front: expected Front to return (4, true), got (%d, %t)", val, ok)
	}

	tmpArr := newQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6, 7}) {
		t.Errorf("Front: expected Values to be [4 5 6 7], got %v", tmpArr)
	}

	newQueue.Clear()
	val, ok = newQueue.Front()
	if ok || val != 0 {
		t.Errorf("Front: expected Front to return (0, false), got (%d, %t)", val, ok)
	}
}

func TestInstance_Enqueue(t *testing.T) {
	newQueue := queue.New()
	newQueue.Enqueue(4, 5, 6)
	tmpArr := newQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6}) {
		t.Errorf("Enqueue 4, 5, 6: expected Values to be [4 5 6], got %v", tmpArr)
	}

	newQueue.Enqueue(7)
	tmpArr = newQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6, 7}) {
		t.Errorf("Enqueue 7: expected Values to be [4 5 6 7], got %v", tmpArr)
	}
}

func TestInstance_Dequeue(t *testing.T) {
	newQueue := queue.New(4, 5, 6)
	newQueue.Enqueue(7)
	val, ok := newQueue.Dequeue()
	if !ok || val != 4 {
		t.Errorf("Dequeue: expected Dequeue to return (4, true), got (%d, %t)", val, ok)
	}

	tmpArr := newQueue.Values()
	if !slicesEqual(tmpArr, []int{5, 6, 7}) {
		t.Errorf("Dequeue: expected Values to be [5 6 7], got %v", tmpArr)
	}

	newQueue.Clear()
	val, ok = newQueue.Dequeue()
	if ok || val != 0 {
		t.Errorf("Dequeue: expected Dequeue to return (0, false), got (%d, %t)", val, ok)
	}
}

func TestInstance_Copy(t *testing.T) {
	newQueue := queue.New(4, 5, 6)
	copiedQueue := newQueue.Copy()
	tmpArr := copiedQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6}) {
		t.Errorf("Copy: expected Values to be [4 5 6], got %v", tmpArr)
	}

	copiedQueue.Enqueue(3)
	tmpArr = copiedQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6, 3}) {
		t.Errorf("Copy: expected Values to be [4 5 6 3], got %v", tmpArr)
	}

	tmpArr = newQueue.Values()
	if !slicesEqual(tmpArr, []int{4, 5, 6}) {
		t.Errorf("Copy: expected Values of original queue to be [4 5 6], got %v", tmpArr)
	}
}

func slicesEqual(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, el := range arr1 {
		if el != arr2[i] {
			return false
		}
	}

	return true
}
