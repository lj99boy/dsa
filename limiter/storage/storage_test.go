package storage

import "testing"

func TestSafeAdd(t *testing.T) {
	if SafeAdd() != 1 {
		t.Error("result wrong")
	}

}
