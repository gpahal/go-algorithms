package set_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/set"
)

func TestNewHashSet(t *testing.T) {
	newSet := set.NewHashSet(4, 5, 6)
	if newSet.Size() != 3 {
		t.Errorf("NewHashSet 4, 5, 6: expected Size to be 3, got %d", newSet.Size())
	}
	if !newSet.Contains(4, 5, 6) {
		t.Error("NewHashSet 4, 5, 6: expected Contains to be true, got false")
	}

	testInterfaceHelper(set.NewHashSet, t)
}
