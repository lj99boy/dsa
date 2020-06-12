//插入排序 就是将一个无序的元素插入一段有序的队列，arr[0]本身已经是有序的
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
