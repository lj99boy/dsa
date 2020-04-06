//插值查找
//是二分查找的优化，比起直接用1/2作为判断的中点，插值查找先计算offset，以这个为基准做比较点
package interpolationsearch

func InterpolationSearch(haystack []int, needle int, low int, high int) int {
	if low < 0 || high >= len(haystack) {
		return -1
	}
	mid := low + (needle-haystack[low])/(haystack[high]-haystack[low])*(high-low)
	if haystack[mid] > needle {
		return InterpolationSearch(haystack, needle, low, mid-1)
	} else if haystack[mid] < needle {
		return InterpolationSearch(haystack, needle, mid+1, high)
	} else {
		return mid
	}
}
