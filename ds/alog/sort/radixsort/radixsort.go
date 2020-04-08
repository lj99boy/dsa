package radixsort

func RadixSort(arr []int) []int {
	//找出最大数为位数基准
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= max {
			max = arr[i]
		}
	}

	//根据max的位数进行，有几位比较几轮
	for ; max > 0; max /= 10 {
		for i := 0; i < len(arr); i++ {

		}
	}
}
