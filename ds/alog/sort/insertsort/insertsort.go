//插入排序
package insertsort

func InsertSort(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := i + 1; j > 0 && j <= len(arr)-1; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
