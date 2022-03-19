package canPartition

func canPartition(nums []int) bool {

	tSum :=0
	for i:=0;i<len(nums);i++{
		tSum+=nums[i]
	}

	if tSum%2!=0{
		return false
	}

	tSum = tSum/2

	return do(tSum,nums,[]int{},len(nums)/2)
}

func do(sum int,nums []int,cNums []int,limit int) bool {


	tSum := 0
	for i:=0;i<len(cNums);i++{
		tSum += cNums[i]
	}

	if tSum == sum {
		return true
	}

	for i:=0;i<len(nums);i++{
		tCNums := append(cNums,nums[i])
		if len(tCNums)>limit{
			return false
		}
		if do(sum,nums[i+1:],tCNums,limit){
			return true
		}
		tCNums = tCNums[:len(tCNums)-1]
	}

	return false
}
