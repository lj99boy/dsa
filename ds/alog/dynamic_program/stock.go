package dynamic_program

import "math"

func MaxProfitOneTrans(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(-prices[i])))
	}

	return dp[len(prices)-1][0]
}

func MaxProfitAnyTimes(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i])))
	}

	return dp[len(prices)-1][0]
}

func MaxProfitAnyTimesWithFee(prices []int, fee int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]-fee
			continue
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i]-fee)))
	}

	return dp[len(prices)-1][0]
}

func MaxProfitDesiredTimes(prices []int) int {
	desiredTimes := 2
	dp := make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, desiredTimes+1)
		for j := 0; j < desiredTimes+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(prices); i++ {
		for j := desiredTimes; j >= 1; j-- {
			if i-1 < 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = int(math.Max(float64(dp[i-1][j][0]), float64(dp[i-1][j][1]+prices[i])))
			dp[i][j][1] = int(math.Max(float64(dp[i-1][j][1]), float64(dp[i-1][j-1][0]-prices[i])))
		}
	}

	return dp[len(prices)-1][desiredTimes][0]
}

func MaxProfitInputDesiredTimes(k int, prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	dp := make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			if i-1 < 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = int(math.Max(float64(dp[i-1][j][0]), float64(dp[i-1][j][1]+prices[i])))
			dp[i][j][1] = int(math.Max(float64(dp[i-1][j][1]), float64(dp[i-1][j-1][0]-prices[i])))
		}
	}

	return dp[len(prices)-1][k][0]
}

func MaxProfitAnyTimesWithCooldown(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		if i-2 < 0 {
			dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[0][0]-prices[i])))
		} else {
			dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-2][0]-prices[i])))
		}
	}

	return dp[len(prices)-1][0]
}
