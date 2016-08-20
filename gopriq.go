// Package gopriq implements a priority queue in Go.
package gopriq

// I'm using my own implementation of a heap.
import (
	"github.com/RMMoreton/goheap"
)

// A Sortable type can be ordered.
type Sortable goheap.Sortable

// A PriQ is just a Heap.
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
