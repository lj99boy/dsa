package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	data := []int{2, 3, 6, 8}
	var res int
	if res = LinearSearch(data, 6); res == -1 {
		t.Error("res wrong")
	}
	println(res)
}
