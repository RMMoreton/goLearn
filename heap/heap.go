//Package heap implements a heap.
package heap

// A Sortable type may be ordered.
type Sortable interface {
	// ComesBefore reports whether the Sortable value which it was called on
	// should come before b.
	ComesBefore(b interface{}) bool
}

// A Heap is just a an array of things that can be ordered.
type Heap []Sortable

// Len returns the length of the Heap.
func (h *Heap) Len() int {
	return len(*h)
}

// Empty reports whether the Heap is empty.
func (h *Heap) Empty() bool {
	if len(*h) == 0 {
		return true
	}
	return false
}

// Add adds a value to the Heap.
func (h *Heap) Add(e Sortable) {
	*h = append(*h, e)
	h.bubbleUp()
}

// Peek returns the element with highest priority without removing it, and
// a boolean to indicate if the returned value is valid, i.e. if there are
// no elements in the queue, the bool will be false.
func (h *Heap) Peek() (Sortable, bool) {
	if len(*h) == 0 {
		return nil, false
	}
	return (*h)[0], true
}

// Remove removes the first element in the Heap, re-heapifies the Heap,
// and returns the removed element. The bool indicates if the returned
// value is valid, similarly to Peek().
func (h *Heap) Remove() (Sortable, bool) {
	// Quick sanity check.
	if len(*h) == 0 {
		return nil, false
	}
	heap := *h
	toReturn := heap[0]
	heap[0] = heap[len(*h)-1]
	*h = heap[:len(*h)-1]
	h.bubbleDown()
	return toReturn, true
}

// Union takes two Heaps and returns a new Heap with all the elements of
// both Heaps. The original Heaps are preserved.
func Union(g, h Heap) Heap {
	var newHeap Heap
	for i := 0; i < len(h); i++ {
		newHeap.Add(h[i])
	}
	for i := 0; i < len(g); i++ {
		newHeap.Add(g[i])
	}
	return newHeap
}

// Heapify makes a Heap using the passed slice. The passed slice is
// preserved.
func Heapify(arr []Sortable) Heap {
	var h Heap
	for i := 0; i < len(arr); i++ {
		h.Add(arr[i])
	}
	return h
}

// bubbleUp shifts the element at the end of the backing array
// up the Heap until the heap property is restored.
func (h *Heap) bubbleUp() {
	i := len(*h) - 1
	// Quick sanity check.
	if i <= 0 {
		return
	}
	heap := *h
	for i > 0 {
		parentI := (i - 1) / 2
		if !heap[i].ComesBefore(heap[parentI]) {
			return
		}
		heap[i], heap[parentI] = heap[parentI], heap[i]
		i = parentI
	}
}

// bubbleDown shifts the 0th element of the backing array
// down the Heap until the heap property is restored.
func (h *Heap) bubbleDown() {
	i := 0
	// Quick sanity check.
	if len(*h) == 0 {
		return
	}
	heap := *h
	for {
		c1 := 2*i + 1
		c2 := 2*i + 2
		// A bunch of edge cases.
		if c1 >= len(*h) && c2 >= len(*h) {
			return
		}
		if c1 >= len(*h) {
			if !heap[i].ComesBefore(heap[c2]) {
				heap[i], heap[c2] = heap[c2], heap[i]
			}
			return
		}
		if c2 >= len(*h) {
			if !heap[i].ComesBefore(heap[c1]) {
				heap[i], heap[c1] = heap[c1], heap[i]
			}
			return
		}
		// Grab the index of the higher priority child.
		var bigger int
		if heap[c1].ComesBefore(heap[c2]) {
			bigger = c1
		} else {
			bigger = c2
		}
		// Compare the bubble-down element's priority with that of
		// it's "biggest" child.
		if heap[i].ComesBefore(heap[bigger]) {
			return
		}
		heap[i], heap[bigger] = heap[bigger], heap[i]
		i = bigger
	}
}
