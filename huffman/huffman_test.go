package huffman

import (
	"github.com/RMMoreton/goLearn/priq"
	"os"
	"testing"
)

var _ priq.PriQ

// TestEncode tests the Encode function.
func TestEncode(t *testing.T) {
	in, err := os.Open("test_in.txt")
	if err != nil {
		t.Errorf("could not open input file")
	}
	defer func() {in.Close()}()
	out, err := os.Create("test_out.txt")
	if err != nil {
		t.Errorf("could not open output file")
	}
	defer func() {out.Close()}()
	Encode(in, out)
}
