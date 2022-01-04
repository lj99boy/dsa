package lengthOfLIS

import "math"

func lengthOfLIS(nums []int) int {
	//tDp := make([]int,len(nums))
	//
	//return doDp(nums,&tDp,)

	dp := make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		dp[i] = 1
	}

	res :=1

	for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}

	return res
}

