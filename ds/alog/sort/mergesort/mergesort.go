//归并排序，就是拆分数组到最小单元，然后从最小单元开始，两两最小单元进行比较排序，胜出的一个较大单元再和其他胜出的较大单元进行排序（两两抽出应该排序）
// 时间复杂度nlogn（因为拆成logn层且每层有至多n个元素比较），空间复杂度n，因为最大占用空间就是n
//https://blog.csdn.net/bjweimengshu/article/details/102384930
package mergesort

func MergeSortRecur(arr []int) []int {
	if len(arr) < 1 {
		return nil
	}

	mid := len(arr) / 2
	if mid < 1 {
		return arr
	}

	//两个排序子集
	left := MergeSortRecur(arr[0:mid])
	right := MergeSortRecur(arr[mid:])

	holder := make([]int, 0)
	//取出两个已经排好序的子集，分别从头拿出来再排序，两两比较小的放前面（从小到大排法）
	for ; len(left) != 0 || len(right) != 0; {
		if len(left) == 0 {
			holder = append(holder, right...)
			break
		}
		if len(right) == 0 {
			holder = append(holder, left...)
			break
		}
		min := 0
		//小的胜出后从数组里拿出来，再循环
		if left[0] > right[0] {
			min = right[0]
			right = right[1:]
		} else {
			min = left[0]
			left = left[1:]
		}
		holder = append(holder, min)
	}
	return holder
}

//func MergeSortIter(arr []int) []int {
//	if len(arr) < 1 {
//		return nil
//	}
//
//	var holder []int
//	for i := 1; i <= len(arr)/2; i *= 2 {
//		for j := 0; j <= len(arr)-1; j += 2 * i {
//			left := arr[j : j+i]
//			right := arr[j+i : j+2*i]
//
//		}
//	}
//}
