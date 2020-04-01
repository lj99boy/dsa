package binarysearch

func BinarySearch(needle int, haystack []int) int {
	low := 0
	high := len(haystack) - 1
	for ; high >= low; {
		mid := (low + high) / 2
		if needle > haystack[mid] {
			low = mid
		} else if needle < haystack[mid] {
			high = mid
		} else if haystack[mid] == needle {
			return mid
		}
	}
	return -1
}
