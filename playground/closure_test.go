package playground

import "testing"

func TestTClosure(t *testing.T) {
	fn := TClosure()

	println(fn())
	println(fn())
	println(fn())
}
