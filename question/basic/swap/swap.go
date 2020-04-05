package swap

func swap(left int, right int) (int, int) {
	left = left - right
	right = right + left
	left = right - left
	return left, right
}
