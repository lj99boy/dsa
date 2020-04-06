//冒泡排序
package bubblesort

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for innerI := 0; innerI < i; innerI++ {
			if arr[innerI] < arr[innerI+1] {
				arr[innerI] = arr[innerI] - arr[innerI+1]
				arr[innerI+1] = arr[innerI] + arr[innerI]
				arr[innerI] = arr[innerI+1] - arr[innerI]
			}
		}
	}
	return arr
}
