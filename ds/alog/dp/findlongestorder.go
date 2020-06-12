package dp

func FindLongestOrder(order []int) (res int) {
	var dp = make([]int, len(order))

	if len(order) < 2 {
		res = len(order)
		return
	}
	res = -1
	dp[0] = 1
	for i := 1; i < len(order); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if order[i] > order[j] {
				dp[i] = returnMax(dp[i], dp[j]+1)
			}
		}
	}

	for _, val := range dp {
		if val > res {
			res = val
		}
	}
	return
}

func returnMax(left int, right int) int {
	if left > right {
		return left
	}
	return right
}
