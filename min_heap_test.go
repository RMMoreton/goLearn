package heap

import (
	"testing"
	"fmt"
)

// Sint stands for sortable int.
type Sint int

// Less allows Sint's to be sorted by the heap functions.
func (a Sint) LessThan(b Interface) bool {
	if int(a) < b.(int) {
		return true
	}
	return false
}

// Test that an empty MinHeap is nil.
func TestEmptyMinHeap(t *testing.T) {
	var h MinHeap
	if nil != h {
		t.Error("h should be nil, but it's", h)
	}
}

// Test creation of a MinHeap, and adding one item.
func TestAddOne(t *testing.T) {
	var h MinHeap
	x := Sint(5)
	h.Add(x)
	expected := "[5]"
	actual := fmt.Sprintf("%v", h)
	if actual != expected {
		t.Error("incorrect output\nexpect:", expected, "\nactual:", actual)
	}
}