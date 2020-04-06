package fibonacci

import (
	"dsa/question/basic"
	"testing"
)

func TestFibonacci(t *testing.T) {
	if basic.Fibonacci(5) != 5 {
		t.Error("res wrong")
	}
}
