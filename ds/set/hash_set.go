package set

type hashSet struct {
	m map[int]struct{}
}

// NewHashSet returns a new hash set with the given elements added to it.
func NewHashSet(items ...int) Interface {
	newSet := &hashSet{
		m: nil,
	}

	newSet.Add(items...)
	return newSet
}

func (hs *hashSet) Size() int {
	return len(hs.m)
}

func (hs *hashSet) Empty() bool {
	return len(hs.m) == 0
}

func (hs *hashSet) Clear() {
	hs.m = nil
}

func (hs *hashSet) Values() []int {
	items := make([]int, len(hs.m))

	idx := 0
	for item := range hs.m {
		items[idx] = item
		idx += 1
	}

	return items
}

func (hs *hashSet) Each(fn func(int) bool) {
	for item := range hs.m {
		if fn(item) {
			break
		}
	}
}

func (hs *hashSet) Iterator() Iterable {
	return &hashSetIterable{
		values:     hs.Values(),
		currentIdx: -1,
	}
}

func (hs *hashSet) Add(items ...int) {
	if hs.m == nil {
		hs.m = make(map[int]struct{}, len(items))
	}

	for _, item := range items {
		hs.m[item] = struct{}{}
	}
}

func (hs *hashSet) Remove(items ...int) {
	if hs.m == nil {
		return
	}

	for _, item := range items {
		delete(hs.m, item)
	}
}

func (hs *hashSet) Contains(items ...int) bool {
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

func (hs *hashSet) Copy() Interface {
	var m map[int]struct{}
	if hs.m != nil {
		m = make(map[int]struct{}, len(hs.m))
		for item := range hs.m {
			m[item] = struct{}{}
		}
	}

	return &hashSet{
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

	hsi.currentIdx += 1
	return true
}

func (hsi *hashSetIterable) Value() int {
	if hsi.currentIdx < 0 || hsi.currentIdx >= len(hsi.values) {
		return 0
	}

	return hsi.values[hsi.currentIdx]
}
