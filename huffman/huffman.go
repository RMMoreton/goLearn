// Package huffman implements Huffman-coding.
package huffman

import (
	"github.com/RMMoreton/goLearn/priq"
	"os"
	"fmt"
	"bufio"
)

var _ priq.PriQ
var _ = fmt.Sprintf("")

// A node is used to build the huffman tree. It has a count, some data, and
// pointers to it's children.
type node struct {
	count int
	data byte
	left *node
	right *node
}

// ComesBefore reports whether the calling node comes before the argument node
// in the priority queue.
func (a *node) ComesBefore(b interface{}) bool {
	bb, ok := b.(*node)
	if !ok {
		// TODO: figure out what to do
		panic("A non-node got into the priority queue")
	}
	if a.count < bb.count {
		return true
	}
	return false
}

// Encode takes an input file and an output file; the input file is Huffman
// coded, and the result is written to the output file.
func Encode(in, out *os.File) {
	counts := count(in)
	p := makePQ(counts)
	h := makeHuffman(p)
	m := make(map[byte]string)
	fillMap(h, m, "")
	for k, v := range m {
		fmt.Printf("k: %c, v: %s\n", k, v)
	}
}

// count takes an input file and counts the number of occurrences of each byte.
// Any reads performed after count returns will read from the beginning of
// the file.
func count(in *os.File) []int {
	in.Seek(0, 0)
	defer func() {in.Seek(0, 0)}()
	counts := make([]int, 256)
	rd := bufio.NewReader(in)
	b, err := rd.ReadByte()
	for err == nil {
		counts[b]++
		b, err = rd.ReadByte()
	}
	return counts
}

// makePQ takes a slice of integers and turns it into a priority queue of nodes.
// Counts of zero are ignored.
func makePQ(counts []int) priq.PriQ {
	var p priq.PriQ
	for i, c := range counts {
		if c == 0 {
			continue
		}
		n := node{c, byte(i), nil, nil}
		p.Add(&n)
	}
	return p
}

// makeHuffman takes a priority queue and turns it into a Huffman tree.
func makeHuffman(p priq.PriQ) *node {
	rr, ok := p.Remove()
	if !ok {
		panic("not enough elements in the priority queue to make a huffman tree")
	}	
	r := rr.(*node)
	ll, ok := p.Remove()
	if !ok {
		panic("not enough elements in the priority queue to make a huffman tree")
	}
	l := ll.(*node)	
	for !p.Empty() {
		parent := new(node)
		parent.count = l.count + r.count
		parent.left = l
		parent.right = r
		p.Add(parent)

		rr, ok = p.Remove()
		r = rr.(*node)
		ll, ok = p.Remove()
		l = ll.(*node)
	}
	root := new(node)
	root.count = l.count + r.count
	root.left = l
	root.right = r
	return root
}

// makeMap takes a huffman tree and returns a map from bytes to strings
// indicating how to compress a given byte.
func fillMap(h *node, m map[byte]string, pre string) {
	if h.left == nil {
		m[h.data] = pre
		return
	}
	fillMap(h.left, m, pre + "1")
	fillMap(h.right, m, pre + "0")
}