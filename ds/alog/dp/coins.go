package dp

// 动态规划一个场景问题，给定硬币面值，找出预定金额最少需要由几枚硬币组成
// 动态规划三要素 1.最优子结构（子问题之间保持独立） 2. 重叠子问题（重复求解子问题可以通过缓存处理） 3. 状态转移方程（类似于在子问题之间转移时的暴力求解方法）
// 自顶向下是递归，自底向上是迭代，也是更标准的动态规划
// https://mp.weixin.qq.com/s/Cw39C9MY9Wr2JlcvBQZMcA
// FindCoinsNum amount是预定金额，coins []int{1,2,5}
func FindCoinsNum(amount int, coins []int) (nums int) {
	//dp表大小应该金额本身(其实问题求解里再大一点也没关系 但是浪费空间了)，因为金额本身也隐含有最多多少个子问题（dp(1),dp(2),...,dp(amount))，这里dp表的含义是dp(5)为金额为5最少需要几个硬币
	dp := make([]int, amount+1)
	for i := 1; i <= len(dp)-1; i++ {
		dp[i] = 998
	}
	dp[0] = 0

	for i := 1; i <= len(dp)-1; i++ {
		for _, coin := range coins {
			tmpVal := 0
			// 这表示负数金额dp(<0)需要的银币数，不存在这种组合，不需要考虑
			if i-coin < 0 {
				continue
			} else {
				tmpVal = dp[i-coin] + 1
			}
			// 当前dp需要的最小硬币数，最小情况下就是各种上一个状态的银币数+1
			dp[i] = min(dp[i], tmpVal)
		}
	}

	return dp[amount]
}

func min(left int, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}
