package set_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/set"
)

func TestNewHashSet(t *testing.T) {
	newSet := set.NewHashSet(4, 5, 6)
	if newSet.Size() != 3 {
		t.Errorf("NewHashSet: expected Size to be 3, got %d", newSet.Size())
	}
	if !newSet.Contains(4, 5, 6) {
		t.Error("NewHashSet: expected Contains to be true, got false")
	}
}
