package isarmstrongnum

import "testing"

func TestIsArmStrongNum(t *testing.T) {
	if !IsArmStrongNum(153) {
		t.Error("wrong res")
	}
}
