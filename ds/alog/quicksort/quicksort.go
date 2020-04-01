package quicksort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(QuickSort(left), append([]int{pivot}, QuickSort(right)...)...)
}
