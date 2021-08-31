package dp

import "math"

func bags01() {
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	maxW := 4
	dp := make([][]int, len(wt)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, maxW+1)
	}

	for i := 1; i <= len(wt); i++ {
		for j := 1; j <= maxW; j++ {
			if j-wt[i] < 0 {
				dp[i][j] = dp[i-1][j]
			}else{
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-wt[i]+val[i]])))
			}
		}
	}
}
