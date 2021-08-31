//快速排序
//取一个标志，小于标志的放一边，大于标志的放另一边，再对这两边递归
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

func QuickSortSingleMap(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	mapPtr := 0
	for i, v := range arr {
		if v < pivot {
			mapPtr++
			arr[mapPtr], arr[i] = arr[i], arr[mapPtr]
		}
	}
	arr[mapPtr], arr[0] = arr[0], arr[mapPtr]

	return append(QuickSortSingleMap(arr[:mapPtr]), append([]int{arr[mapPtr]}, QuickSortSingleMap(arr[mapPtr+1:])...)...)
}
