package euclidean_alog

func euclideanAlog(left int, right int) int {
	if left == right {
		return left
	}
	if left < right {
		left, right = right, left
	}

	remainder := left % right
	if remainder == 1 {
		return remainder
	}
	if right%remainder == 0 {
		return remainder
	}

	return euclideanAlog(right, remainder)
}
