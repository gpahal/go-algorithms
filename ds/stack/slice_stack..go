package stack

// SliceStack represents a stack instance implemented as a slice used ascircular buffer.
type SliceStack struct {
	arr   []int
	start int
	size  int
}

// NewSliceStack returns a new slice stack instance with the given items pushed into it.
func NewSliceStack(items ...int) Interface {
	return newSliceStack(items...)
}

func newSliceStack(items ...int) *SliceStack {
	s := &SliceStack{}
	s.Push(items...)
	return s
}

// Len returns the number of items in the stack.
func (s *SliceStack) Len() int {
	return s.size
}

// Empty checks whether the stack is empty.
func (s *SliceStack) Empty() bool {
	return s.size == 0
}

// Clear deletes all the items from the stack.
func (s *SliceStack) Clear() {
	s.arr = nil
	s.start = 0
	s.size = 0
}

// Top returns the top/last pushed element of the stack. If the stack is empty, second return
// value is false.
func (s *SliceStack) Top() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	return s.arr[(s.start+s.size-1)%len(s.arr)], true
}

// Push pushes the given items to the stack.
func (s *SliceStack) Push(items ...int) {
	if len(items) == 0 {
		return
	}

	for i := 0; i < len(s.arr)-s.size && i < len(items); i++ {
		s.arr[(s.start+s.size+i)%len(s.arr)] = items[i]
	}
	if len(items) <= len(s.arr)-s.size {
		s.size += len(items)
		return
	}

	// update items and s.size
	items = items[len(s.arr)-s.size:]
	s.size = len(s.arr)

	// slice extension needed
	if s.start == 0 {
		s.arr = append(s.arr, items...)
		s.size = len(s.arr)
		return
	}

	s.arr = append(s.arr, items...)
	j := s.size
	for i := 0; i < s.start; i++ {
		s.arr[j%len(s.arr)] = s.arr[i]
		j++
	}
	for i := 0; i < len(items); i++ {
		s.arr[j%len(s.arr)] = items[i]
		j++
	}
	s.size = len(s.arr)
}

// Pop pops out an item from the stack in LIFO (Last In First Out) order. If the stack is empty,
// second return value is false.
func (s *SliceStack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	v := s.arr[(s.start+s.size-1)%len(s.arr)]
	s.size--
	return v, true
}

// Copy creates a new copy of the stack.
func (s *SliceStack) Copy() Interface {
	newStack := newSliceStack()
	if s.size == 0 {
		return newStack
	}

	newStack.start = 0
	newStack.size = s.size
	newStack.arr = make([]int, s.size)
	if s.start+s.size > len(s.arr) {
		copy(newStack.arr[:len(s.arr)-s.start], s.arr[s.start:len(s.arr)])
		copy(newStack.arr[len(s.arr)-s.start:], s.arr[:s.start+s.size-len(s.arr)])
	} else {
		copy(newStack.arr, s.arr[s.start:])
	}

	return newStack
}
