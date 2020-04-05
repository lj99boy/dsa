package regular

import "testing"

func TestInts_Iterator(t *testing.T) {
	ints := Ints{3, 2, 5}

	it := ints.Iterator()

	for v, ok := it(); !ok; {
		print(v)
	}
}
