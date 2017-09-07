package list

type DoublyLinkedList struct {
	head *Element
	tail *Element
}

func NewDoublyLinkedList(items ...int) Interface {
	newList := &DoublyLinkedList{}
	newList.PushBack(items...)
	return newList
}

func (dll *DoublyLinkedList) Len() int {
	if dll.head == nil {
		return 0
	}

	count := 1
	for curr := dll.head; curr.Next != nil; curr = curr.Next {
		count++
	}

	return count
}

func (dll *DoublyLinkedList) Empty() bool {
	return dll.head == nil
}

func (dll *DoublyLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
}

func (dll *DoublyLinkedList) Values() []int {
	if dll.head == nil {
		return []int{}
	}

	items := make([]int, 0)
	for curr := dll.head; curr != nil; curr = curr.Next {
		items = append(items, curr.Value)
	}

	return items
}

func (dll *DoublyLinkedList) Each(fn func(int) bool) {
	if dll.head == nil {
		return
	}

	for curr := dll.head; curr != nil; curr = curr.Next {
		if fn(curr.Value) {
			break
		}
	}
}

func (dll *DoublyLinkedList) Iterator() Iterable {
	return &doublyLinkedListIterable{
		curr: dll.head,
	}
}

func (dll *DoublyLinkedList) First() *Element {
	return dll.head
}

func (dll *DoublyLinkedList) Last() *Element {
	return dll.tail
}

func (dll *DoublyLinkedList) At(i int) *Element {
	if dll.head == nil {
		return nil
	}

	reverse := false
	if i < 0 {
		reverse = true
		i = (i + 1) * -1
	}

	currIdx := 0
	var curr *Element
	if reverse {
		curr = dll.tail
	} else {
		curr = dll.head
	}

	for curr != nil {
		if currIdx == i {
			return curr
		}

		currIdx++
		if reverse {
			curr = curr.Prev
		} else {
			curr = curr.Next
		}
	}

	return nil
}

func (dll *DoublyLinkedList) PushFront(items ...int) {
	if len(items) == 0 {
		return
	}

	var next *Element
	var curr *Element
	for _, item := range items {
		if curr == nil {
			curr = &Element{
				Value: item,
				Next:  dll.head,
			}

			if dll.head != nil {
				dll.head.Prev = curr
			} else {
				dll.tail = curr
			}
		} else {
			curr = &Element{
				Value: item,
				Next:  curr,
			}

			next.Prev = curr
		}

		next = curr
	}

	dll.head = curr
}

func (dll *DoublyLinkedList) PopFront() *Element {
	if dll.head == nil {
		return nil
	}

	removedEl := dll.head
	dll.head = dll.head.Next
	if dll.head == nil {
		dll.tail = nil
	} else {
		dll.head.Prev = nil
	}

	removedEl.Next = nil
	return removedEl
}

func (dll *DoublyLinkedList) PushBack(items ...int) {
	if len(items) == 0 {
		return
	}

	var prev *Element
	var curr *Element
	for _, item := range items {
		if curr == nil {
			curr = &Element{
				Value: item,
				Prev:  dll.tail,
			}

			if dll.tail != nil {
				dll.tail.Next = curr
			} else {
				dll.head = curr
			}
		} else {
			curr = &Element{
				Value: item,
				Prev:  curr,
			}

			prev.Next = curr
		}

		prev = curr
	}

	dll.tail = curr
}

func (dll *DoublyLinkedList) PopBack() *Element {
	if dll.head == nil {
		return nil
	}

	removedEl := dll.tail
	dll.tail = dll.tail.Prev
	if dll.tail == nil {
		dll.head = nil
	} else {
		dll.tail.Next = nil
	}

	removedEl.Prev = nil
	return removedEl
}

func (dll *DoublyLinkedList) InsertAt(i int, item int) *Element {
	if dll.head == nil {
		if i == 0 || i == -1 {
			dll.PushFront(item)
			return dll.First()
		}

		return nil
	}
	if i == 0 {
		return dll.InsertAfter(dll.head, item)
	}
	if i > 0 {
		i--
	}

	curr := dll.At(i)
	return dll.InsertAfter(curr, item)
}

func (dll *DoublyLinkedList) InsertAfter(e *Element, item int) *Element {
	if e == nil {
		return nil
	}

	newEl := &Element{
		Value: item,
		Next:  e.Next,
		Prev:  e,
	}
	if e.Next == nil {
		dll.tail = newEl
	} else {
		e.Next.Prev = newEl
	}
	e.Next = newEl

	return newEl
}

func (dll *DoublyLinkedList) InsertBefore(e *Element, item int) *Element {
	if e == nil {
		return nil
	}

	newEl := &Element{
		Value: item,
		Next:  e,
		Prev:  e.Prev,
	}
	if e.Prev == nil {
		dll.head = newEl
	} else {
		e.Prev.Next = newEl
	}
	e.Prev = newEl

	return newEl
}

func (dll *DoublyLinkedList) RemoveAt(i int) *Element {
	return dll.Remove(dll.At(i))
}

func (dll *DoublyLinkedList) Remove(e *Element) *Element {
	if e == nil {
		return nil
	}

	if dll.head == e {
		dll.head = e.Next
	}
	if dll.tail == e {
		dll.tail = e.Prev
	}

	if e.Prev != nil {
		e.Prev.Next = e.Next
	}
	if e.Next != nil {
		e.Next.Prev = e.Prev
	}
	e.Prev = nil
	e.Next = nil

	return e
}

func (dll *DoublyLinkedList) RemoveAfter(e *Element) *Element {
	if e == nil {
		return nil
	}

	return dll.Remove(e.Next)
}

func (dll *DoublyLinkedList) RemoveBefore(e *Element) *Element {
	if e == nil {
		return nil
	}

	return dll.Remove(e.Prev)
}

func (dll *DoublyLinkedList) DeleteFirst(items ...int) {
	if len(items) == 0 || dll.head == nil {
		return
	}

	itemsMap := make(map[int]struct{}, len(items))
	for _, item := range items {
		itemsMap[item] = struct{}{}
	}

	var next *Element
	for curr := dll.head; curr != nil; curr = next {
		next = curr.Next
		if _, ok := itemsMap[curr.Value]; ok {
			delete(itemsMap, curr.Value)
			dll.Remove(curr)

			if len(itemsMap) == 0 {
				break
			}
		}
	}
}

func (dll *DoublyLinkedList) Delete(items ...int) {
	if len(items) == 0 || dll.head == nil {
		return
	}

	itemsMap := make(map[int]struct{}, len(items))
	for _, item := range items {
		itemsMap[item] = struct{}{}
	}

	var next *Element
	for curr := dll.head; curr != nil; curr = next {
		next = curr.Next
		if _, ok := itemsMap[curr.Value]; ok {
			dll.Remove(curr)
		}
	}
}

func (dll *DoublyLinkedList) Contains(items ...int) bool {
	if len(items) == 0 {
		return true
	}
	if dll.head == nil {
		return false
	}

	itemsMap := make(map[int]struct{}, len(items))
	for _, item := range items {
		itemsMap[item] = struct{}{}
	}

	for curr := dll.head; curr != nil; curr = curr.Next {
		if _, ok := itemsMap[curr.Value]; ok {
			delete(itemsMap, curr.Value)
		}
	}

	return len(itemsMap) == 0
}

func (dll *DoublyLinkedList) Copy() Interface {
	return NewDoublyLinkedList(dll.Values()...)
}

type doublyLinkedListIterable struct {
	curr  *Element
	value int
}

func (dlli *doublyLinkedListIterable) Next() bool {
	if dlli.curr == nil {
		dlli.value = 0
		return false
	}

	dlli.value = dlli.curr.Value
	dlli.curr = dlli.curr.Next
	return true
}

func (dlli *doublyLinkedListIterable) Value() int {
	return dlli.value
}
