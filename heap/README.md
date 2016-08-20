# goLearn/heap

heap provides a Golang implementation of a heap using
slices as the backing data structure.

[Documentation](http://godoc.org/github.com/RMMoreton/goLearn/heap)

## Examples

### Sortable Interface

	type Sint int

	func (a Sint) ComesBefore(b interface{}) bool {
		// Type assert b back to an Sint.
		bb, ok := b.(Sint)
		if !ok {
			// Do something to save yourself, or just
			panic("AHHHH!!!!")
		}
		// In this example, higher value means higher priority.
		if int(a) > int(bb) {
			return true
		}
		return false
	}

### Creating a Heap

	var h heap.Heap

The zero-value of a Heap is (happily) useable without any extra initialization.