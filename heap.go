/*
 * Package heap contains implementations of both a min-heap and a max-heap
 */
package heap

import (

)

// Interface is used to make my code prettier.
type Interface interface {}

// Sortable values may be sorted.
type Sortable interface {
	// ComesFirst reports whether the Sortable value which it was called on
	// should come before b.
	ComesFirst(b Interface) bool
}

// MinHeap is a minimum-heap.
type MinHeap []Sortable

// Add adds a value to the min-heap.
// TODO: make it shift up to the correct position.
func (h *MinHeap) Add(e Sortable) {
	*h = append(*h, e)
}