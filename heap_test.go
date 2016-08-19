package heap

import (
	"fmt"
	"testing"
)

// Sint stands for sortable int.
type Sint int

// Less allows Sint's to be sorted by the heap functions.
func (a Sint) ComesBefore(iB Interface) bool {
	b := iB.(Sint)
	if a.Val() < b.Val() {
		return true
	}
	return false
}

// Val gives me a programatic way to access an Sint's value,
// which is useful in ComesBefore().
func (a Sint) Val() int {
	return int(a)
}

// Test that an empty Heap is nil.
func TestEmptyHeap(t *testing.T) {
	var h Heap
	if nil != h {
		t.Error("h should be nil, but it's", h)
	}
}

// Test that an empty Heap has length 0.
func TestEmptyHeapLength(t *testing.T) {
	var h Heap
	if 0 != h.Len() {
		t.Error("empty heap should have length 0")
	}
}

// Test creation of a Heap, and adding one item.
func TestAddOne(t *testing.T) {
	var h Heap
	x := Sint(5)
	h.Add(x)
	expected := "[5]"
	actual := fmt.Sprintf("%v", h)
	if actual != expected {
		t.Error("incorrect output\nexpect:", expected, "\nactual:", actual)
	}
}

// Test that Adding increases length.
func TestAddIncreasesLength(t *testing.T) {
	var h Heap
	h.Add(Sint(1))
	if 1 != h.Len() {
		t.Error("adding an element should increase the length")
	}
}

// Test adding two items, where the second should not replace
// the first.
func TestAddTwoNoBubble(t *testing.T) {
	var h Heap
	x := Sint(5)
	y := Sint(6)
	h.Add(x)
	h.Add(y)
	expected := "[5 6]"
	actual := fmt.Sprintf("%v", h)
	if actual != expected {
		t.Error("incorrect output\nexpect:", expected, "\nactual:", actual)
	}
}

// Test adding two items, where the second should replace the first.
func TestAddTwoYesBubble(t *testing.T) {
	var h Heap
	x := Sint(6)
	y := Sint(5)
	h.Add(x)
	h.Add(y)
	expected := "[5 6]"
	actual := fmt.Sprintf("%v", h)
	if actual != expected {
		t.Error("incorrect output\nexpect:", expected, "\nactual:", actual)
	}
}

// Test adding 16 items, where they should all bubble up to the top.
func TestAdd16IncreasingPriority(t *testing.T) {
	var h Heap
	for i := 15; i >= 0; i-- {
		h.Add(Sint(i))
		p := h.Peek().(Sint)
		if i != p.Val() {
			t.Error("adding elements in increasing priority\nexpect:", i, "\nactual:", p.Val())
		}
	}
}

// Test Peeking into a nil heap.
func TestPeekIntoNil(t *testing.T) {
	var h Heap
	p := h.Peek()
	if nil != p {
		t.Error("peeking into a nil heap should return nil")
	}
}

// Test Peeking into an empty (but non-nil) heap.
func TestPeekEmptyNonNil(t *testing.T) {
	var h Heap
	h.Add(Sint(5))
	_ = h.Remove()
	p := h.Peek()
	if nil != p {
		t.Error("peeking into an empty, non-nil heap should return nil")
	}
}

// Test Peeking into a non-nil heap.
func TestPeekNonNil(t *testing.T) {
	var h Heap
	for i := 15; i > 0; i-- {
		h.Add(Sint(i))
		p := h.Peek().(Sint)
		if p.Val() != i {
			t.Error("peeking into a heap\nexpect:", i, "\nactual:", p.Val())
		}
	}
}

// Test removing from a nil heap.
func TestRemoveNil(t *testing.T) {
	var h Heap
	r := h.Remove()
	if nil != r {
		t.Error("removing from nil heap should return nil")
	}
}

// Test removing from a non-nil heap.
func TestRemoveNonNil(t *testing.T) {
	var h Heap
	for i := 15; i >= 8; i-- {
		h.Add(Sint(i))
	}
	p := h.Remove().(Sint)
	if 8 != p.Val() {
		t.Error("removing from a non-nil heap\nexpect:", 8, "\nactual:", p.Val())
	}
}

// Test that removing from a heap actually makes the length
// of the heap go down.
func TestRemoveDecreaseLength(t *testing.T) {
	var h Heap
	for i := 7; i >= 0; i-- {
		h.Add(Sint(i))
	}
	_ = h.Remove()
	if 7 != h.Len() {
		t.Error("after removal, length should be 7")
	}
}

// Test that removing bubbles down correctly.
func TestRemoveBubbleDown(t *testing.T) {
	var h Heap
	for i := 15; i >= 0; i-- {
		h.Add(Sint(i))
	}
	for i := 0; i <= 15; i++ {
		p := h.Remove().(Sint)
		if i != p.Val() {
			t.Error("removing doesn't preserve the heap property\nexpect:", i, "actual:", p.Val())
		}
	}
}
