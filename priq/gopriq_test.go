package gopriq

import (
	"testing"
)

// Sint stands for sortable int.
type Sint int

// ComesBefore reports whether the calling value comes before the called
// value.
func (a Sint) ComesBefore(b interface{}) bool {
	bb, ok := b.(Sint)
	if !ok {
		panic("AAHHHHH!")
	}
	if int(a) < int(bb) {
		return true
	}
	return false
}

// Test Empty on an empty Priority Queue.
func TestEmptyOnEmptyPriQ(t *testing.T) {
	var q PriQ
	if !q.Empty() {
		t.Error("priority queue should be empty")
	}
}

// Test Empty on a non-empty Priority Queue.
func TestEmptyOnNonEmptyPriQ(t *testing.T) {
	var q PriQ
	q.Add(Sint(5))
	if q.Empty() {
		t.Error("priority queue should not be empty")
	}
}

// Test that Adding two elements results in correct order.
func TestAddTwoNeedReorder(t *testing.T) {
	var q PriQ
	q.Add(Sint(5))
	q.Add(Sint(4))
	e, ok := q.Remove()
	if !ok {
		t.Error("should be able to remove from non-empty PQ")
	}
	ee, ok := e.(Sint)
	if !ok {
		t.Error("non-Sint value got in to the PQ")
	}
	if int(ee) != 4 {
		t.Error("priority queue returns results in the wrong order")
	}
	e, ok = q.Remove()
	if !ok {
		t.Error("should be able to remove from non-empty PQ")
	}
	ee, ok = e.(Sint)
	if !ok {
		t.Error("non-Sint value got in to the PQ")
	}
	if int(ee) != 5 {
		t.Error("priority queue returns results in the wrong order")
	}
}

// Test that I can't remove from a zero-value PQ.
func TestRemovalFromNil(t *testing.T) {
	var q PriQ
	_, ok := q.Remove()
	if ok {
		t.Error("should not be able to remove from an empty queue")
	}
}

// Test that I can't remove from an empty PQ.
func TestRemovalFromEmpty(t *testing.T) {
	var q PriQ
	q.Add(Sint(5))
	_, ok := q.Remove()
	if !ok {
		t.Error("should be able to remove from non-empty queue")
	}
	_, ok = q.Remove()
	if ok {
		t.Error("should not be able to remove from non-empty queue")
	}
}
