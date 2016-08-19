//Package heap contains implementations of both a min-heap and a max-heap
package heap

// Interface is used to make my code prettier.
type Interface interface{}

// Sortable values may be sorted.
type Sortable interface {
	// ComesBefore reports whether the Sortable value which it was called on
	// should come before b.
	ComesBefore(b Interface) bool
}

// Heap is just a an array of things that can be ordered.
type Heap []Sortable

// Len returns the length of the heap.
func (h *Heap) Len() int {
	return len(*h)
}

// Add adds a value to the heap.
func (h *Heap) Add(e Sortable) {
	*h = append(*h, e)
	h.bubbleUp()
}

// Peek returns the element with highest priority without removing it.
func (h *Heap) Peek() (Interface, bool) {
	if len(*h) == 0 {
		return nil, false
	}
	return (*h)[0], true
}

// Remove removes the first element in the heap, re-heapifies the heap,
// and returns the removed element.
func (h *Heap) Remove() (Interface, bool) {
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

// Union takes two heaps and returns a new Heap with the elements of
// both. It's on the user to make sure that the two heaps are of the
// same type.
func (h *Heap) Union(g Heap) Heap {
	var newHeap Heap
	for i := 0; i < len(*h); i++ {
		newHeap.Add((*h)[i])
	}
	for i := 0; i < len(g); i++ {
		newHeap.Add(g[i])
	}
	return newHeap
}

// Heapify makes a heap using the passed slice. The passed slice is
// preserved.
func Heapify(arr []Sortable) Heap {
	var h Heap
	for i := 0; i < len(arr); i++ {
		h.Add(arr[i])
	}
	return h
}

// bubbleUp shifts the element at the end of the backing array
// up the heap until the heap property is restored.
func (h *Heap) bubbleUp() {
	i := len(*h) - 1
	// Quick sanity check.
	if i <= 0 {
		return
	}
	heap := *h
	for {
		parentI := (i - 1) / 2
		if !heap[i].ComesBefore(heap[parentI]) {
			return
		}
		heap[i], heap[parentI] = heap[parentI], heap[i]
		i = parentI
	}
}

// bubbleDown shifts the 0th element of the backing array
// down the heap until the heap property is restored.
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
