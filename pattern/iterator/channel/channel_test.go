package channel

import "testing"

func TestInts1_Iterator(t *testing.T) {
	tmpInts1 := Ints{3, 2, 1}
	for v := range tmpInts1.Iterator() {
		println(v)
	}
}
