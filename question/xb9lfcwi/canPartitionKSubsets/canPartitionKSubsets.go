package canPartitionKSubsets

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {

	if k > len(nums) {
		return false
	}

	sum := 0 
	
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	
	if sum % k != 0 {
		return false
	}

	buckets := make([]int,k)

	sort.Ints(nums)

	for i,j:=0,len(nums)-1;i<j;{
		nums[i],nums[j] = nums[j],nums[i]
		i++
		j--
	}

	return recur(nums,0,buckets,sum/k)
}

func recur(nums []int,nIndex int,buckets []int,target int) bool {

	if nIndex == len(nums) {
		for i:=0;i<len(buckets);i++{
			if buckets[i]!=target{
				return false
			}
		}
		return true
	}

	for i:=0;i<len(buckets);i++{
		if buckets[i]+nums[nIndex] > target{
			continue
		}

		buckets[i] += nums[nIndex]
		if recur(nums,nIndex+1,buckets,target) {
			return true
		}
		buckets[i] -= nums[nIndex]
	}

	return false
}
