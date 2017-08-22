package set

type HashSet struct {
	m map[int]struct{}
}

// NewHashSet returns a new hash set with the given elements added to it.
func NewHashSet(items ...int) Interface {
	newSet := &HashSet{
		m: nil,
	}

	newSet.Add(items...)
	return newSet
}

func (hs *HashSet) Size() int {
	return len(hs.m)
}

func (hs *HashSet) Empty() bool {
	return len(hs.m) == 0
}

func (hs *HashSet) Clear() {
	hs.m = nil
}

func (hs *HashSet) Values() []int {
	items := make([]int, len(hs.m))

	idx := 0
	for item := range hs.m {
		items[idx] = item
		idx++
	}

	return items
}

func (hs *HashSet) Each(fn func(int) bool) {
	for item := range hs.m {
		if fn(item) {
			break
		}
	}
}

func (hs *HashSet) Iterator() Iterable {
	return &hashSetIterable{
		values:     hs.Values(),
		currentIdx: -1,
	}
}

func (hs *HashSet) Add(items ...int) {
	if hs.m == nil {
		hs.m = make(map[int]struct{}, len(items))
	}

	for _, item := range items {
		hs.m[item] = struct{}{}
	}
}

func (hs *HashSet) Remove(items ...int) {
	if hs.m == nil {
		return
	}

	for _, item := range items {
		delete(hs.m, item)
	}
}

func (hs *HashSet) Contains(items ...int) bool {
	if hs.m == nil {
		return len(items) == 0
	}

	var ok bool
	for _, item := range items {
		_, ok = hs.m[item]
		if !ok {
			return false
		}
	}

	return true
}

func (hs *HashSet) Copy() Interface {
	var m map[int]struct{}
	if hs.m != nil {
		m = make(map[int]struct{}, len(hs.m))
		for item := range hs.m {
			m[item] = struct{}{}
		}
	}

	return &HashSet{
		m: m,
	}
}

type hashSetIterable struct {
	values     []int
	currentIdx int
}

func (hsi *hashSetIterable) Next() bool {
	if hsi.currentIdx >= len(hsi.values)-1 {
		return false
	}

	hsi.currentIdx++
	return true
}

func (hsi *hashSetIterable) Value() int {
	if hsi.currentIdx < 0 || hsi.currentIdx >= len(hsi.values) {
		return 0
	}

	return hsi.values[hsi.currentIdx]
}
