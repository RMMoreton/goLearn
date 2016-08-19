# Golang-Heap

Golang-Heap provides a Golang implementation of a heap using
slices as the backing data structure.

[Documentation](http://godoc.org/github.com/RMMoreton/golang-heap)

## Examples

### Sortable Interface

	type Sint int

	func (a Sint) ComesBefore(b interface{}) bool {
		// Type assert b back to an Sint.
		bb, ok := b.(Sint)
		if !ok {
			// Do something to save yourself
		}
		// In this example, higher value means higher priority.
		if int(a) > int(bb) {
			return true
		}
		return false
	}

Obviously you can get yourself in to trouble with the type assert; be
careful about what you put into the Heap and you'll be okay.

### Creating a Heap

	var h goheap.Heap

The zero-value of a Heap is useable without any extra initialization.