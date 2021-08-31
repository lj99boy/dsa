package other

import "sort"

func FourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	var threeSum = func(innerNums []int, tmpT int) [][]int {
		res := make([][]int, 0)
		length := len(innerNums)
		for i := 0; i < length; i++ {
			if i > 0 && innerNums[i] == innerNums[i-1] {
				continue
			}
			offset := tmpT - innerNums[i]
			l := i + 1
			r := length - 1
			for ; l < r; {
				if l > i+1 && innerNums[l] == innerNums[l-1] {
					l++
					continue
				}
				if r < len(innerNums)-2 && innerNums[r] == innerNums[r+1] {
					r--
					continue
				}

				if offset > (innerNums[l] + innerNums[r]) {
					l++
				} else if offset < (innerNums[l] + innerNums[r]) {
					r--
				} else {
					res = append(res, []int{innerNums[i], innerNums[l], innerNums[r]})
					l++
				}
			}
		}
		return res
	}
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		offset := target - nums[i]

		tmpRes := threeSum(nums[i+1:], offset)
		for k := range tmpRes {
			tmpRes[k] = append(tmpRes[k], nums[i])
		}
		ans = append(ans, tmpRes...)
	}

	return ans
}

//func threeSum(nums []int) [][]int {
//	res := make([][]int, 0)
//	sort.Ints(nums)
//	length := len(nums)
//	for i := 0; i < length; i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//
//		offset := 0 - nums[i]
//		l := i + 1
//		r := length - 1
//		for ; l < r; {
//			if l > i+1 && nums[l] == nums[l-1] {
//				l++
//				continue
//			}
//			if r < len(nums)-2 && nums[r] == nums[r+1] {
//				r--
//				continue
//			}
//
//			if offset > (nums[l] + nums[r]) {
//				l++
//			} else if offset < (nums[l] + nums[r]) {
//				r--
//			} else {
//				res = append(res, []int{nums[i], nums[l], nums[r]})
//				l++
//			}
//		}
//	}
//
//	return res
//}
