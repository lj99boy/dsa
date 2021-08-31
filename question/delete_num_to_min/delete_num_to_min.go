package delete_num_to_min

// 给一个整数，删掉几位后让它保持最小
// 其实这道题直接想还是有点难，但是考虑只删一位的情况，然后递归到好几位，就简单了
// 只删一位的时候从最左看起，如果右边大于左边，不能删当前左边，继续往后，直到出现左边大于右边，表示删了左边，右边替代左边会让当前数变小

func deleteNumToMin(arr []int, nums int) []int {

	if nums > 0 {
		nums--
	} else {
		return arr
	}

	return deleteNumToMin(pickMinArr(arr), nums)
}

func pickMinArr(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return append(arr[:i], arr[i+1:]...)
		}
	}

	return arr[:len(arr)-1]
}
