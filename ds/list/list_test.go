package list_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/list"
)

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

func testInterfaceHelper(newFn func(...int) list.Interface, t *testing.T) {
	l := newFn()
	if l.Len() != 0 {
		t.Errorf("Len: expected Len to be 0, got %d", l.Len())
	}
	if !l.Empty() {
		t.Error("Empty: expected Empty to be true, got false")
	}

	emptyList := newFn()

	l.PushFront(9, 7, 5)
	tmpArr := l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9}) {
		t.Errorf("PushFront 9, 7, 5: expected Values to be [5 7 9], got %v", tmpArr)
	}

	el := l.PopFront()
	if el == nil || el.Value != 5 {
		t.Errorf("PopFront: expected PopFront value to be 5, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{7, 9}) {
		t.Errorf("PopFront: expected Values to be [7 9], got %v", tmpArr)
	}
	l.PushFront(5)

	el = emptyList.PopFront()
	if el != nil {
		t.Errorf("PopFront: expected PopFront to be nil, got %v", el)
	}

	l.PushBack(11, 5, -4)
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5, -4}) {
		t.Errorf("PushBack 11, 5, -4: expected Values to be [5 7 9 11 5 -4], got %v", tmpArr)
	}

	el = l.PopBack()
	if el == nil || el.Value != -4 {
		t.Errorf("PopBack: expected PopBack value to be -4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5}) {
		t.Errorf("PopBack: expected Values to be [5 7 9 11 5], got %v", tmpArr)
	}
	l.PushBack(-4)

	el = emptyList.PopBack()
	if el != nil {
		t.Errorf("PopBack: expected PopBack to be nil, got %v", el)
	}

	el = l.First()
	if el == nil || el.Value != 5 {
		t.Errorf("First: expected First value to be 5, got %v", el)
	}

	el = emptyList.First()
	if el != nil {
		t.Errorf("First: expected First to be nil, got %v", el)
	}

	el = l.Last()
	if el == nil || el.Value != -4 {
		t.Errorf("Last: expected Last value to be -4, got %v", el)
	}

	el = emptyList.Last()
	if el != nil {
		t.Errorf("Last: expected Last to be nil, got %v", el)
	}

	el = l.At(2)
	if el == nil || el.Value != 9 {
		t.Errorf("At 2: expected At value to be 9, got %v", el)
	}

	el = l.At(-2)
	if el == nil || el.Value != 5 {
		t.Errorf("At -2: expected At value to be 5, got %v", el)
	}

	el = l.At(9)
	if el != nil {
		t.Errorf("At 9: expected At to be nil, got %v", el)
	}

	el = l.InsertAt(2, 4)
	if el == nil || el.Value != 4 {
		t.Errorf("InsertAt 2, 4: expected InsertAt value to be 4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 4, 9, 11, 5, -4}) {
		t.Errorf("InsertAt 2, 4: expected Values to be [5 7 4 9 11 5 -4], got %v", tmpArr)
	}

	el = l.InsertAt(10, 4)
	if el != nil {
		t.Errorf("InsertAt 10, 4: expected InsertAt to be nil, got %v", el)
	}

	el = l.RemoveAt(2)
	if el == nil || el.Value != 4 {
		t.Errorf("RemoveAt 2: expected RemoveAt value to be 4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5, -4}) {
		t.Errorf("RemoveAt 2: expected Values to be [5 7 9 11 5 -4], got %v", tmpArr)
	}

	el = l.RemoveAt(10)
	if el != nil {
		t.Errorf("RemoveAt 10: expected RemoveAt to be nil, got %v", el)
	}

	el = l.At(2)
	el = l.Remove(el)
	if el == nil || el.Value != 9 {
		t.Errorf("Remove el: expected Remove value to be 9, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 11, 5, -4}) {
		t.Errorf("Remove el: expected Values to be [5 7 11 5 -4], got %v", tmpArr)
	}

	el = l.InsertAt(2, 9)
	if el == nil || el.Value != 9 {
		t.Errorf("InsertAt 2, 9: expected InsertAt value to be 9, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5, -4}) {
		t.Errorf("InsertAt 2, 9: expected Values to be [5 7 9 11 5 -4], got %v", tmpArr)
	}

	el = l.InsertAfter(el, 4)
	if el == nil || el.Value != 4 {
		t.Errorf("InsertAfter el, 4: expected InsertAfter value to be 4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 4, 11, 5, -4}) {
		t.Errorf("InsertAfter el, 4: expected Values to be [5 7 4 9 11 5 -4], got %v", tmpArr)
	}

	el = l.InsertAfter(nil, 4)
	if el != nil {
		t.Errorf("InsertAfter nil, 4: expected InsertAfter to be nil, got %v", el)
	}

	el = l.At(2)
	el = l.RemoveAfter(el)
	if el == nil || el.Value != 4 {
		t.Errorf("RemoveAfter el: expected RemoveAfter value to be 4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5, -4}) {
		t.Errorf("RemoveAfter el: expected Values to be [5 7 9 11 5 -4], got %v", tmpArr)
	}

	el = l.RemoveAfter(nil)
	if el != nil {
		t.Errorf("RemoveAfter nil: expected RemoveAfter to be nil, got %v", el)
	}

	el = l.At(2)
	el = l.InsertBefore(el, 4)
	if el == nil || el.Value != 4 {
		t.Errorf("InsertBefore el, 4: expected InsertBefore value to be 4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 4, 9, 11, 5, -4}) {
		t.Errorf("InsertBefore el, 4: expected Values to be [5 7 4 9 11 5 -4], got %v", tmpArr)
	}

	el = l.InsertBefore(nil, 4)
	if el != nil {
		t.Errorf("InsertBefore nil, 4: expected InsertBefore to be nil, got %v", el)
	}

	el = l.At(3)
	el = l.RemoveBefore(el)
	if el == nil || el.Value != 4 {
		t.Errorf("RemoveBefore el: expected RemoveBefore value to be 4, got %v", el)
	}
	tmpArr = l.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5, -4}) {
		t.Errorf("RemoveBefore el: expected Values to be [5 7 9 11 5 -4], got %v", tmpArr)
	}

	el = l.RemoveBefore(nil)
	if el != nil {
		t.Errorf("RemoveBefore nil: expected RemoveBefore to be nil, got %v", el)
	}

	newList := l.Copy()
	tmpArr = newList.Values()
	if !slicesEqual(tmpArr, []int{5, 7, 9, 11, 5, -4}) {
		t.Errorf("Copy: expected Values to be [5 7 9 11 5 -4], got %v", tmpArr)
	}

	newList.DeleteFirst(7, 5)
	tmpArr = newList.Values()
	if !slicesEqual(tmpArr, []int{9, 11, 5, -4}) {
		t.Errorf("DeleteFirst: expected Values to be [9 11 5 -4], got %v", tmpArr)
	}

	newList.PushFront(7, 5)

	newList.Delete(7, 5)
	tmpArr = newList.Values()
	if !slicesEqual(tmpArr, []int{9, 11, -4}) {
		t.Errorf("Delete: expected Values to be [9 11 -4], got %v", tmpArr)
	}

	if !newList.Contains(-4, 9) {
		t.Errorf("Contains -4, 9: expected Contains to be true, got false %v", l.Values())
	}

	if newList.Contains(-4, 5) {
		t.Error("Contains -4, 5: expected Contains to be false, got true")
	}

	tmpArr = make([]int, 0)
	newList.Each(func(item int) bool {
		if !newList.Contains(item) {
			t.Errorf("Each 9, 11, -4 (Contains %d): expected Contains to be true, got false", item)
			return true
		}

		tmpArr = append(tmpArr, item)
		return false
	})
	if !slicesEqual(tmpArr, []int{9, 11, -4}) {
		t.Errorf("Each 9, 11, -4: expected iterations to be [9 11 -4], got %v", tmpArr)
	}

	tmpArr = make([]int, 0)
	iterable := newList.Iterator()
	for iterable.Next() {
		item := iterable.Value()
		if !newList.Contains(item) {
			t.Errorf("Each 9, 11, -4 (Contains %d): expected Contains to be true, got false", item)
			break
		}

		tmpArr = append(tmpArr, item)
	}
	if !slicesEqual(tmpArr, []int{9, 11, -4}) {
		t.Errorf("Each 9, 11, -4: expected iterations to be [9 11 -4], got %v", tmpArr)
	}

	l.Clear()
	if l.Len() != 0 {
		t.Errorf("Clear: expected Len to be 0, got %d", l.Len())
	}
	if !l.Empty() {
		t.Error("Clear: expected Empty to be true, got false")
	}
}
