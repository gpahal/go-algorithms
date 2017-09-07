package stack_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/stack"
)

func TestNew(t *testing.T) {
	newStack := stack.New(4, 5, 6)
	if newStack.Len() != 3 {
		t.Errorf("New 4, 5, 6: expected Len to be 3, got %d", newStack.Len())
	}

	tmpArr := newStack.Values()
	if !slicesEqual(tmpArr, []int{6, 5, 4}) {
		t.Errorf("New 4, 5, 6: expected Values to be [6 5 4], got %v", tmpArr)
	}
}

func TestInstance_Len(t *testing.T) {
	newStack := stack.New()
	if newStack.Len() != 0 {
		t.Errorf("Len: expected Len to be 0, got %d", newStack.Len())
	}

	newStack.Push(4, 5, 6)
	if newStack.Len() != 3 {
		t.Errorf("Len: expected Len to be 3, got %d", newStack.Len())
	}
}

func TestInstance_Empty(t *testing.T) {
	newStack := stack.New()
	if !newStack.Empty() {
		t.Error("Empty: expected Empty to be true, got false")
	}

	newStack.Push(4, 5, 6)
	if newStack.Empty() {
		t.Error("Empty: expected Empty to be false, got true")
	}
}

func TestInstance_Clear(t *testing.T) {
	newStack := stack.New(4, 5, 6)
	newStack.Clear()
	if newStack.Len() != 0 {
		t.Errorf("Clear: expected Len to be 0, got %d", newStack.Len())
	}
}

func TestInstance_Values(t *testing.T) {
	newStack := stack.New()
	tmpArr := newStack.Values()
	if len(tmpArr) != 0 {
		t.Errorf("Values: expected Values to be [], got %#v", tmpArr)
	}

	newStack.Push(4, 5, 6)
	tmpArr = newStack.Values()
	if !slicesEqual(tmpArr, []int{6, 5, 4}) {
		t.Errorf("Values: expected Values to be [6 5 4], got %v", tmpArr)
	}
}

func TestInstance_Top(t *testing.T) {
	newStack := stack.New(4, 5, 6)
	newStack.Push(7)
	val, ok := newStack.Top()
	if !ok || val != 7 {
		t.Errorf("Top: expected Top to return (7, true), got (%d, %t)", val, ok)
	}

	tmpArr := newStack.Values()
	if !slicesEqual(tmpArr, []int{7, 6, 5, 4}) {
		t.Errorf("Top: expected Values to be [7 6 5 4], got %v", tmpArr)
	}

	newStack.Clear()
	val, ok = newStack.Top()
	if ok || val != 0 {
		t.Errorf("Top: expected Top to return (0, false), got (%d, %t)", val, ok)
	}
}

func TestInstance_Push(t *testing.T) {
	newStack := stack.New(4, 5, 6)
	newStack.Push(7)
	tmpArr := newStack.Values()
	if !slicesEqual(tmpArr, []int{7, 6, 5, 4}) {
		t.Errorf("Push 7: expected Values to be [7 6 5 4], got %v", tmpArr)
	}
}

func TestInstance_Pop(t *testing.T) {
	newStack := stack.New(4, 5, 6)
	newStack.Push(7)
	val, ok := newStack.Pop()
	if !ok || val != 7 {
		t.Errorf("Pop: expected Pop to return (7, true), got (%d, %t)", val, ok)
	}

	tmpArr := newStack.Values()
	if !slicesEqual(tmpArr, []int{6, 5, 4}) {
		t.Errorf("Pop: expected Values to be [6 5 4], got %v", tmpArr)
	}

	newStack.Clear()
	val, ok = newStack.Pop()
	if ok || val != 0 {
		t.Errorf("Pop: expected Pop to return (0, false), got (%d, %t)", val, ok)
	}
}

func TestInstance_Copy(t *testing.T) {
	newStack := stack.New(4, 5, 6)
	copiedStack := newStack.Copy()
	tmpArr := copiedStack.Values()
	if !slicesEqual(tmpArr, []int{6, 5, 4}) {
		t.Errorf("Copy: expected Values to be [6 5 4], got %v", tmpArr)
	}

	copiedStack.Push(3)
	tmpArr = copiedStack.Values()
	if !slicesEqual(tmpArr, []int{3, 6, 5, 4}) {
		t.Errorf("Copy: expected Values to be [3 6 5 4], got %v", tmpArr)
	}

	tmpArr = newStack.Values()
	if !slicesEqual(tmpArr, []int{6, 5, 4}) {
		t.Errorf("Copy: expected Values of original stack to be [6 5 4], got %v", tmpArr)
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
