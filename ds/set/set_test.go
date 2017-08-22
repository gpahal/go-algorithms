package set_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/set"
)

func sliceContains(arr []int, items ...int) bool {
outer:
	for _, item := range items {
		for _, el := range arr {
			if el == item {
				continue outer
			}
		}

		return false
	}

	return true
}

func TestInterface(t *testing.T) {
	newSet := set.NewHashSet()
	if newSet.Size() != 0 {
		t.Errorf("Size: expected Size to be 0, got %d", newSet.Size())
	}
	if !newSet.Empty() {
		t.Error("Empty: expected Empty to be true, got false")
	}

	newSet.Add(5, 7, 9)
	newSet.Add(11, 5, -4)
	if newSet.Size() != 5 {
		t.Errorf("Add 5, 7, 9, 11, -4: expected Size to be 5, got %d", newSet.Size())
	}

	newSet.Delete(5, 7, 13)
	if newSet.Size() != 3 {
		t.Errorf("Delete 5, 7, 13: expected Size to be 3, got %d", newSet.Size())
	}

	if !newSet.Contains(-4, 9) {
		t.Errorf("Contains -4, 9: expected Contains to be true, got false %v", newSet.Values())
	}

	if newSet.Contains(-4, 5) {
		t.Error("Contains -4, 5: expected Contains to be false, got true")
	}

	tmpSet := newSet.Copy()
	if !set.AreEqual(newSet, tmpSet) {
		t.Error("Copy: expected AreEqual to be true, got false")
	}
	tmpSet.Clear()
	if set.AreEqual(newSet, tmpSet) {
		t.Error("Clear Copy: expected AreEqual to be false, got true")
	}

	tmpArr := newSet.Values()
	if len(tmpArr) != 3 {
		t.Errorf("Values 9, 11, -4: expected list's length to be 3, got %d", len(tmpArr))
	}
	if !sliceContains(tmpArr, 9, 11, -4) {
		t.Errorf("Values 9, 11, -4: expected list to contain [9 11 -4], got %v", tmpArr)
	}

	tmpArr = make([]int, 0)
	newSet.Each(func(item int) bool {
		if !newSet.Contains(item) {
			t.Errorf("Each 9, 11, -4 (Contains %d): expected Contains to be true, got false", item)
			return true
		}

		tmpArr = append(tmpArr, item)
		return false
	})
	if len(tmpArr) != 3 {
		t.Errorf("Each 9, 11, -4: expected total iterations to be 3, got %d", len(tmpArr))
	}
	if !sliceContains(tmpArr, 9, 11, -4) {
		t.Errorf("Each 9, 11, -4: expected iterations to contain [9 11 -4], got %v", tmpArr)
	}

	tmpArr = make([]int, 0)
	iterable := newSet.Iterator()
	for iterable.Next() {
		item := iterable.Value()
		if !newSet.Contains(item) {
			t.Errorf("Each 9, 11, -4 (Contains %d): expected Contains to be true, got false", item)
			break
		}

		tmpArr = append(tmpArr, item)
	}
	if len(tmpArr) != 3 {
		t.Errorf("Each 9, 11, -4: expected total iterations to be 3, got %d", len(tmpArr))
	}
	if !sliceContains(tmpArr, 9, 11, -4) {
		t.Errorf("Each 9, 11, -4: expected iterations to contain [9 11 -4], got %v", tmpArr)
	}

	newSet.Clear()
	if newSet.Size() != 0 {
		t.Errorf("Clear: expected Size to be 0, got %d", newSet.Size())
	}
	if !newSet.Empty() {
		t.Error("Clear: expected Empty to be true, got false")
	}
}

func TestAreEqual(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)

	if !set.AreEqual(newSet1, newSet3) {
		t.Error("AreEqual: expected true, got false")
	}
	if set.AreEqual(newSet1, newSet2) {
		t.Error("AreEqual: expected false, got true")
	}
}

func TestIsSubset(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet4 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4, 90, 12)

	if !set.IsSubset(newSet1, newSet4) {
		t.Error("IsSubset: expected true, got false")
	}
	if !set.IsSubset(newSet1, newSet3) {
		t.Error("IsSubset: expected true, got false")
	}
	if set.IsSubset(newSet1, newSet2) {
		t.Error("IsSubset: expected false, got true")
	}
	if set.IsSubset(newSet4, newSet1) {
		t.Error("IsSubset: expected false, got true")
	}
}

func TestIsSuperset(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet4 := set.NewHashSet(1, 7, 8, 10, 12)

	if !set.IsSuperset(newSet1, newSet4) {
		t.Error("IsSuperset: expected true, got false")
	}
	if !set.IsSuperset(newSet1, newSet3) {
		t.Error("IsSuperset: expected true, got false")
	}
	if set.IsSuperset(newSet1, newSet2) {
		t.Error("IsSuperset: expected false, got true")
	}
	if set.IsSuperset(newSet4, newSet1) {
		t.Error("IsSuperset: expected false, got true")
	}
}

func TestAreDisjoint(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewHashSet(-1, -7, -8, -10, -12, -15, 4)

	if !set.AreDisjoint(newSet1, newSet3) {
		t.Error("AreDisjoint: expected true, got false")
	}
	if set.AreDisjoint(newSet1, newSet2) {
		t.Error("AreDisjoint: expected false, got true")
	}
}

func TestMergeInto(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedMergedSet := set.NewHashSet(1, 7, 8, 10, 12, 15, -4, 3, -9, -2, 18)

	set.MergeInto(newSet1, newSet2)
	if !set.AreEqual(expectedMergedSet, newSet1) {
		t.Errorf("MergeInto: expected %v, got %v", expectedMergedSet.Values(), newSet1.Values())
	}
}

func TestRetainOnly(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedRetainedSet := set.NewHashSet(7, 8, 12, 15)

	set.RetainOnly(newSet1, newSet2)
	if !set.AreEqual(expectedRetainedSet, newSet1) {
		t.Errorf("RetainOnly: expected %v, got %v", expectedRetainedSet.Values(), newSet1.Values())
	}
}

func TestSeparateFrom(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedSeparatedSet := set.NewHashSet(1, 10, -4)

	set.SeparateFrom(newSet1, newSet2)
	if !set.AreEqual(expectedSeparatedSet, newSet1) {
		t.Errorf("SeparateFrom: expected %v, got %v", expectedSeparatedSet.Values(), newSet1.Values())
	}
}

func TestUnion(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedUnionSet := set.NewHashSet(1, 7, 8, 10, 12, 15, -4, 3, -9, -2, 18)

	gotUnionSet := set.Union(newSet1, newSet2)
	if !set.AreEqual(expectedUnionSet, gotUnionSet) {
		t.Errorf("Union: expected %v, got %v", expectedUnionSet.Values(), gotUnionSet.Values())
	}
}

func TestIntersection(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedIntersectionSet := set.NewHashSet(7, 8, 12, 15)

	gotIntersectionSet := set.Intersection(newSet1, newSet2)
	if !set.AreEqual(expectedIntersectionSet, gotIntersectionSet) {
		t.Errorf("Intersection: expected %v, got %v", expectedIntersectionSet.Values(), gotIntersectionSet.Values())
	}
}

func TestDifference(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedDifferenceSet := set.NewHashSet(1, 10, -4)

	gotDifferenceSet := set.Difference(newSet1, newSet2)
	if !set.AreEqual(expectedDifferenceSet, gotDifferenceSet) {
		t.Errorf("Difference: expected %v, got %v", expectedDifferenceSet.Values(), gotDifferenceSet.Values())
	}
}

func TestSymmetricDifference(t *testing.T) {
	newSet1 := set.NewHashSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewHashSet(3, -9, 7, 15, -2, 18, 12, 8)
	expectedSymmetricDifferenceSet := set.NewHashSet(1, 10, -4, 3, -9, -2, 18)

	gotDifferenceSet := set.SymmetricDifference(newSet1, newSet2)
	if !set.AreEqual(expectedSymmetricDifferenceSet, gotDifferenceSet) {
		t.Errorf("Symmetric Difference: expected %v, got %v", expectedSymmetricDifferenceSet.Values(), gotDifferenceSet.Values())
	}
}
