package fibonacci

//Fibonacci 斐波那契数列
func Fibonacci(pos int) int {
	if pos < 0 {
		return -1
	}
	if pos == 0 {
		return 0
	}

	if pos == 1 {
		return 1
	}

	before, after := 0, 1

	for i := 1; i < pos; i++ {
		before, after = after, before+after
	}

	return after
}
