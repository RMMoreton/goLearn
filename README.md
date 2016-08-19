# Golang-Heap

Golang-Heap provides a Golang implementation of a heap using
slices as the backing data structure.

## The Sortable Interface
To implement the Sortable interface, a type must have a method 
`ComesBefore(b interface{}) bool` defined on it. `ComesBefore()` should
return true if the calling value comes before (i.e. has a higher priority
than) `b`, and false otherwise.

Note that because `b` is of type `interface{}`, you will almost certainly
have to type assert it inside your function. So don't be careless about
what you put in your Heap!

## Creation
Create a Heap with the syntax

	var h golang-heap.Heap

The zero value of a Heap is useable without any initialization.