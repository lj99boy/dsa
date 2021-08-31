// 找到最接近且大于当前数列的数列
package find_at_least_max_arr

import "sort"

//重排当前数列找出最靠近当前数列且更大的
//如 12537 ---> 12357
// 12342 -- > 12423
// 这道题的思想是，顺序排列的数列绝对是最小的，逆序绝对是最大的。
// 1. 从右边数起，找到逆序和顺序的交接点
//如 12537 交界点是3,12342交界点是3，
// 2. 然后用交界点替换逆序区域里最靠近且大于的数，替换交界点
// 3. 然后再对逆序区域顺序排序

func findAtLeastMaxArr(arr []int) []int {
	index := findTransferPoint(arr)
	if index == 0 {
		return arr
	}
	nearstKey := findNearstKey(index-1, arr)
	arr[index-1], arr[nearstKey] = arr[nearstKey], arr[index-1]
	sort.Ints(arr[index:])
	return arr
}

func findTransferPoint(arr []int) int {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			return i
		}
	}
	return 0
}

func findNearstKey(transferKey int, arr []int) int {
	nearst := arr[transferKey]
	init := true
	key := 0
	for i := len(arr) - 1; i > transferKey; i-- {
		if init {
			if arr[i] > arr[transferKey] {
				nearst = arr[i]
				init = false
				key = i
			}
		} else {
			if arr[i] > arr[transferKey] && arr[i] < nearst {
				nearst = arr[i]
				key = i
			}
		}
	}
	return key
}
