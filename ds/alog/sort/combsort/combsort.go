//梳排序
//确定一个衰减率，从排序长度开始除以这个衰减率（之后长度再除以衰减率直至整数除法结果为1）作为取值排序的步长，以步长为跨度取值进行排序
//https://blog.csdn.net/disappear_xuechao/article/details/77842675
package combsort

func CombSort(arr []int) []int {
	ratio := 1.3

	for step := int(float64(len(arr)) / ratio); step >= 1; step = int(float64(step) / ratio) {
		//这里要小于step的原因是 step这个位置的元素其实已经被第0位用来比较过了，因为刚好是一个衰减率位置上面的元素
		for i := 0; i < step; i++ {
			for j := i; j <= len(arr)-1; j += step {
				if j+step > len(arr)-1 {
					break
				}
				if arr[j] > arr[j+step] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
				}
			}
		}
	}

	return arr
}
