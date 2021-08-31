package rabin_karp

import "testing"

func TestIsContain(t *testing.T) {
	haystack := "miaomiao123"
	needle := "omi"

	println(IsContain(haystack, needle))

}
