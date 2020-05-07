package dp

import "testing"

func TestRegex(t *testing.T) {
	str := []rune{'a','b','c'}
	//pattern := []rune{'.','b','c'}
	pattern := []rune{'a','.','*'}

	println(Regex(str,pattern))
}
