package set

type Interface interface {
	Size() int
	Empty() bool
	Clear()
	Values() []int
	Each(fn func(int) bool)
	Iterator() Iterable
	Add(items ...int)
	Remove(items ...int)
	Contains(items ...int) bool
	Copy() Interface
}

type Iterable interface {
	Next() bool
	Value() int
}

func AreEqual(set1, set2 Interface) bool {
	if set1.Size() != set2.Size() {
		return false
	}

	equal := true
	set1.Each(func(item int) bool {
		if !set2.Contains(item) {
			equal = false
			return true
		}

		return false
	})

	return equal
}

func IsSubset(subSet, superSet Interface) bool {
	if subSet.Size() > superSet.Size() {
		return false
	}

	subset := true
	subSet.Each(func(item int) bool {
		if !superSet.Contains(item) {
			subset = false
			return true
		}

		return false
	})

	return subset
}

func IsSuperset(superSet, subSet Interface) bool {
	return IsSubset(subSet, superSet)
}

func AreDisjoint(set1, set2 Interface) bool {
	if set1.Size() == 0 {
		return true
	}

	disjoint := true
	set1.Each(func(item int) bool {
		if set2.Contains(item) {
			disjoint = false
			return true
		}

		return false
	})

	return disjoint
}

func MergeInto(mainSet, otherSet Interface) {
	mainSet.Add(otherSet.Values()...)
}

func RetainOnly(mainSet, otherSet Interface) {
	mainSet.Each(func(item int) bool {
		if !otherSet.Contains(item) {
			mainSet.Remove(item)
		}

		return false
	})
}

func SeparateFrom(mainSet, otherSet Interface) {
	mainSet.Remove(otherSet.Values()...)
}

func Union(sets ...Interface) Interface {
	newSet := NewHashSet()
	for _, set := range sets {
		MergeInto(newSet, set)
	}

	return newSet
}

func Intersection(sets ...Interface) Interface {
	if len(sets) == 0 {
		return NewHashSet()
	}

	newSet := sets[0].Copy()
	sets = sets[1:]
	for _, set := range sets {
		RetainOnly(newSet, set)
	}

	return newSet
}

func Difference(sets ...Interface) Interface {
	if len(sets) == 0 {
		return NewHashSet()
	}

	newSet := sets[0].Copy()
	sets = sets[1:]
	for _, set := range sets {
		SeparateFrom(newSet, set)
	}

	return newSet
}

func SymmetricDifference(set1, set2 Interface) Interface {
	return Union(Difference(set1, set2), Difference(set2, set1))
}
