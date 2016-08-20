// Package gopryq implements a priority queue in Go.
package gopriq

// I'm using my own implementation of a heap.
import (
	"github.com/RMMoreton/goheap"
)

// This *might* cause some problems down the road, with some functions expecting
// a goheap.Sortable and getting a gopriq.Sortable. Hopefully not!
type Sortable goheap.Sortable

// A priority queue is just a structure holding a Heap. I tried to write
// `type PriQ goheap.Heap` but I couldn't get the type converting right in
// the functions.
type PriQ struct {
	h goheap.Heap
}

// Add adds a value to the Priority queue.
func (p *PriQ) Add(e Sortable) {
	p.h.Add(e)
}

// Empty reports whether the Priority Queue is empty.
func (p *PriQ) Empty() bool {
	return p.h.Empty()
}

// Remove removes a value from the Priority Queue and returns it. Remove
// also returns a boolean to indicate if the returned value is valid.
func (p *PriQ) Remove() (Sortable, bool) {
	return p.h.Remove()
}
