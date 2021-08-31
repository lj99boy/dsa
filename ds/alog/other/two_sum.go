package other

func twoSum(nums []int, target int) []int {
	hashArr := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashArr[nums[i]] = i
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		offset := target - nums[i]
		if _, ok := hashArr[offset]; ok && hashArr[offset] != i {
			return append(append(res, i), hashArr[offset])
		}
	}

	return []int{-1, -1}
}
