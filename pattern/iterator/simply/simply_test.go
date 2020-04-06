package simply

import "testing"

func TestInts_Do(t *testing.T) {
	ints := Ints{3, 7, 9}
	execFunc := func(v int) {
		print(v)
	}
	ints.Do(execFunc)
}

func TestInts_DoWithEscape(t *testing.T) {
	ints := Ints{3, 7, 9}

	execFunc := func(v int) bool {
		if v == 7 {
			return true
		}
		print(v)
		return false
	}
	ints.DoWithEscape(execFunc)
}
