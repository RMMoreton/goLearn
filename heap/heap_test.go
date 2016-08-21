package heap

import (
	"testing"
	"fmt"
)

// SintNeq stands for sortable int with a non-equals ComesBefore
// method.
type SintNeq int

// SintEq will have a slightly different ComesBefore method defined on it.
type SintEq int

// testCase holds testing cases and their results.
type testCase struct {
	toAdd []int
	expect string
}

// ComesBefore allows SintNeq's to be sorted by the heap functions. Note the strict
// less than.
func (a SintNeq) ComesBefore(b interface{}) bool {
	bb := b.(SintNeq)
	return int(a) < int(bb)
}

// ComesBefore allows SintEq's to be sorted by the heap functions. Note the
// non-strict less than.
func (a SintEq) ComesBefore(b interface{}) bool {
	bb := b.(SintEq)
	return int(a) <= int(bb)
}

// errorf simplifies all the functions that will fail due to having a
// different output than expected.
func errorf(t *testing.T, msg, expect, actual string) {
	s := fmt.Sprintf("%s\nExpect: %s\nActual: %s", msg, expect, actual)
	t.Errorf(s)
}

// runCases runs test cases and compares the output to the expected output,
// logging any errors.
func runCases(t *testing.T, cases []testCase, msg string) {
	for _, c := range cases {
		var hNeq, hEq Heap
		for _, n := range c.toAdd {
			hNeq.Add(SintNeq(n))
			hEq.Add(SintEq(n))
		}
		neqStr := fmt.Sprintf("%v", hNeq)
		eqStr := fmt.Sprintf("%v", hEq)
		if neqStr != c.expect {
			errorf(t, msg, c.expect, neqStr)
		}
		if eqStr != c.expect {
			errorf(t, msg, c.expect, eqStr)
		}
	}
}

// TestLenNilHeap tests the length of a nil Heap.
func TestLenNil(t *testing.T) {
	var h Heap
	l := h.Len()
	if l != 0 {
		t.Errorf("Nil heap has non-zero length.")
	}
}

// TestAddOne tests adding one element to a Heap.
func TestAddOne(t *testing.T) {
	cases := []testCase {
		{
			[]int{5,},
			"[5]",
		},
		{
			[]int{0,},
			"[0]",
		},
		{
			[]int{-5,},
			"[-5]",
		},
	}
	runCases(t, cases, "Adding a single element:")
}

// TestAddTwo tests adding two elements to a Heap.
func TestAddTwo(t *testing.T) {
	cases := []testCase {
		{
			[]int{5, 6,},
			"[5 6]",
		},
		{
			[]int{5, 5,},
			"[5 5]",
		},
		{
			[]int{5, 0,},
			"[0 5]",
		},
		{
			[]int{-1, 0,},
			"[-1 0]",
		},
	}
	runCases(t, cases, "Adding two elements:")
}