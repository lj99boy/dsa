package searchRange

func SearchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1

	left2 := 0
	right2 := len(nums) - 1

	for ; left <= right; {
		mid := left + (right-left)/2

		if target == nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	for ; left2 <= right2; {
		mid := left2 + (right2-left2)/2

		if target == nums[mid] {
			left2 = mid + 1
		} else if target > nums[mid] {
			left2 = mid + 1
		} else if target < nums[mid] {
			right2 = mid - 1
		}
	}
	if left >= len(nums){
		return []int{-1, -1}
	} else if nums[left] == target {
		return []int{left, right2}
	} else {
		return []int{-1, -1}
	}
}
