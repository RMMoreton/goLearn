package heap

import (
	"fmt"
	"math/rand"
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
		pTemp, ok := h.Peek()
		if !ok {
			t.Error("unable to peek into what should be a non-empty Heap")
		}
		p := pTemp.(Sint)
		if i != p.Val() {
			t.Error("adding elements in increasing priority\nexpect:", i, "\nactual:", p.Val())
		}
	}
}

// Test Peeking into a nil heap.
func TestPeekIntoNil(t *testing.T) {
	var h Heap
	_, ok := h.Peek()
	if ok {
		t.Error("was able to peek into what should be a nil Heap")
	}
}

// Test Peeking into an empty (but non-nil) heap.
func TestPeekEmptyNonNil(t *testing.T) {
	var h Heap
	h.Add(Sint(5))
	_, ok := h.Remove()
	if !ok {
		t.Error("should be able to remove from a non-empty Heap")
	}
	_, ok = h.Peek()
	if ok {
		t.Error("was able to peek into what should be an empty Heap")
	}
}

// Test Peeking into a non-nil heap.
func TestPeekNonNil(t *testing.T) {
	var h Heap
	for i := 15; i > 0; i-- {
		h.Add(Sint(i))
		pTemp, ok := h.Peek()
		if !ok {
			t.Error("was unable to peek")
		}
		p := pTemp.(Sint)
		if p.Val() != i {
			t.Error("peeking into a heap\nexpect:", i, "\nactual:", p.Val())
		}
	}
}

// Test removing from a nil heap.
func TestRemoveNil(t *testing.T) {
	var h Heap
	_, ok := h.Remove()
	if ok {
		t.Error("should not be able to remove from a nil Heap")
	}
}

// Test removing from a non-nil heap.
func TestRemoveNonNil(t *testing.T) {
	var h Heap
	for i := 15; i >= 8; i-- {
		h.Add(Sint(i))
	}
	rTemp, ok := h.Remove()
	if !ok {
		t.Error("was not able to remove from a non-empty Heap")
	}
	r := rTemp.(Sint)
	if 8 != r.Val() {
		t.Error("removing from a non-nil heap\nexpect:", 8, "\nactual:", r.Val())
	}
}

// Test that removing from a heap actually makes the length
// of the heap go down.
func TestRemoveDecreaseLength(t *testing.T) {
	var h Heap
	for i := 7; i >= 0; i-- {
		h.Add(Sint(i))
	}
	_, ok := h.Remove()
	if !ok {
		t.Error("should be able to remove from a non-empty Heap")
	}
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
		rTemp, ok := h.Remove()
		if !ok {
			t.Error("should be able to remove from a non-empty Heap")
		}
		r := rTemp.(Sint)
		if i != r.Val() {
			t.Error("removing doesn't preserve the heap property\nexpect:", i, "actual:", r.Val())
		}
	}
}

// Test that a large number of insertions returns them in the
// correct order.
func TestManyInsertions(t *testing.T) {
	var h Heap
	numToAdd := 4096
	for i := 0; i < numToAdd; i++ {
		h.Add(Sint(rand.Int()))
	}
	prevTmp, ok := h.Remove()
	if !ok {
		t.Error("should be able to remove from a non-empty Heap")
	}
	prev := prevTmp.(Sint)
	for i := 1; i < numToAdd; i++ {
		curTmp, ok := h.Remove()
		if !ok {
			t.Error("should be able to remove from a non-empty Heap")
		}
		cur := curTmp.(Sint)
		if cur.Val() < prev.Val() {
			t.Error("elements came out in the wrong order\npre:", prev.Val(), "\ncur:", cur.Val())
		}
		prev = cur
	}
}

// Test that after Heapify, the original slice is unchanged.
func TestHeapifySliceDoesntChange(t *testing.T) {
	var s []Sortable
	for i := 0; i < 16; i++ {
		s = append(s, Sint(rand.Int()))
	}
	expect := fmt.Sprintf("%v", s)
	_ = Heapify(s)
	actual := fmt.Sprintf("%v", s)
	if expect != actual {
		t.Error("heapify changed underlying slice\nexpect:", expect, "\nactual:", actual)
	}
}

// Test that Heapify returns something with the heap property.
func TestHeapifyReturnsHeap(t *testing.T) {
	var s []Sortable
	for i := 0; i < 1024; i++ {
		s = append(s, Sint(rand.Int()))
	}
	h := Heapify(s)
	prevTmp, ok := h.Remove()
	if !ok {
		t.Error("should be able to remove from a non-empty heap")
	}
	prev := prevTmp.(Sint)
	for i := 1; i < 1024; i++ {
		curTmp, ok := h.Remove()
		if !ok {
			t.Error("should be able to remove from a non-empty heap")
		}
		cur := curTmp.(Sint)
		if prev.Val() > cur.Val() {
			t.Error("out of order elements\nfirst:", prev.Val(), "\nsecond:", cur.Val())
		}
		prev = cur
	}
}

// Test that a Union returns an array of the correct length.
func TestUnionHasCorrectLength(t *testing.T) {
	var g, h Heap
	for i := 0; i < 16; i++ {
		g.Add(Sint(rand.Int()))
		h.Add(Sint(rand.Int()))
	}
	u := g.Union(h)
	if 32 != u.Len() {
		t.Error("union'd heap has incorrect length")
	}
}

// Test that a Union doesn't change either Heap.
func TestUnionPreservesBothHeaps(t *testing.T) {
	var g, h Heap
	for i := 0; i < 16; i++ {
		g.Add(Sint(rand.Int()))
		h.Add(Sint(rand.Int()))
	}
	gExpect := fmt.Sprintf("%v", g)
	hExpect := fmt.Sprintf("%v", h)
	_ = g.Union(h)
	gActual := fmt.Sprintf("%v", g)
	hActual := fmt.Sprintf("%v", h)
	if gExpect != gActual {
		t.Error("union changed an underlying Heap\nexpect:", gExpect, "\nactual:", gActual)
	}
	if hExpect != hActual {
		t.Error("union changed an underlying Heap\nexpect:", hExpect, "\nactual:", hActual)
	}
}

// Test that removing from a Union doesn't change either old Heap.
func TestUnionRemovalDoesntAffectOtherHeaps(t *testing.T) {
	var g, h Heap
	for i := 0; i < 16; i++ {
		g.Add(Sint(rand.Int()))
		h.Add(Sint(rand.Int()))
	}
	gExpect := fmt.Sprintf("%v", g)
	hExpect := fmt.Sprintf("%v", h)
	u := g.Union(h)
	_, ok := u.Remove()
	if !ok {
		t.Error("should be able to remove from a non-empty heap")
	}
	gActual := fmt.Sprintf("%v", g)
	hActual := fmt.Sprintf("%v", h)
	if gExpect != gActual {
		t.Error("union removal changed an underlying Heap\nexpect:", gExpect, "\nactual:", gActual)
	}
	if hExpect != hActual {
		t.Error("union removal changed an underlying Heap\nexpect:", hExpect, "\nactual:", hActual)
	}
}

// Test that a Union has all the right elements (in the right order).
func TestUnionRemovalOrderAndContent(t *testing.T) {
	var g, h Heap
	for i := 0; i < 16; i++ {
		g.Add(Sint(rand.Int()))
		h.Add(Sint(rand.Int()))
	}
	u := g.Union(h)
	prevTmp, ok := u.Remove()
	if !ok {
		t.Errorf("should be able to remove from non-empty heap")
	}
	prev := prevTmp.(Sint)
	for i := 1; i < 32; i++ {
		curTmp, ok := u.Remove()
		if !ok {
			t.Errorf("should be able to remove from non-empty heap")
		}
		cur := curTmp.(Sint)
		if prev > cur {
			t.Error("union doesn't yield a heap\nfirst:", prev.Val(), "second:", cur.Val())
		}
	}
}

// Test that Union is the same in either direction.
func TestUnionOrderDoesntMatter(t *testing.T) {
	var g, h Heap
	for i := 0; i < 16; i++ {
		g.Add(Sint(rand.Int()))
		h.Add(Sint(rand.Int()))
	}
	u1 := g.Union(h)
	u2 := h.Union(g)
	for i := 0; i < 32; i++ {
		e1Tmp, ok := u1.Remove()
		if !ok {
			t.Errorf("should be able to remove from non-empty heap")
		}
		e1 := e1Tmp.(Sint)
		e2Tmp, ok := u2.Remove()
		if !ok {
			t.Errorf("should be able to remove from non-empty heap")
		}
		e2 := e2Tmp.(Sint)
		if e1.Val() != e2.Val() {
			t.Error("union order matters\nfirst:", e1.Val(), "\nsecond:", e2.Val())
		}
	}
}

// Test that Union with one nil heap returns correctly.
func TestUnionSingleNil(t *testing.T) {
	var g, h Heap
	for i := 0; i < 16; i++ {
		g.Add(Sint(rand.Int()))
	}
	expect := fmt.Sprintf("%v", g)
	u1 := g.Union(h)
	u2 := h.Union(g)
	actual1 := fmt.Sprintf("%v", u1)
	actual2 := fmt.Sprintf("%v", u2)
	if actual1 != expect {
		t.Error("union with nil\nexpect:", expect, "\nactual:", actual1)
	}
	if actual2 != expect {
		t.Error("union with nil\nexpect:", expect, "\nactual:", actual2)
	}
}

// Test that Union with two nil heaps returns nil.
func TestUnionDoubleNil(t *testing.T) {
	var g, h Heap
	u1 := g.Union(h)
	u2 := g.Union(g)
	if nil != u1 {
		t.Error("union with two nils should be nil, got", u1)
	}
	if nil != u2 {
		t.Error("union with two nils should be nil, got", u2)
	}
}
