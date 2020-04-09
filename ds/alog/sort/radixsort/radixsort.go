package radixsort

import "strconv"

func RadixSort(arr []int) []int {
	//找出最大数为位数基准
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= max {
			max = arr[i]
		}
	}

	//max有几位
	length := getNumLength(max)
	//根据max的位数进行，有几位进行几轮计数排序
	for i := length; i > 0; i-- {
		//十进制整数就十种取值
		holder := make([]int, 10)
		for j := 0; j < len(arr); j++ {
			holder[getNumByPosition(arr[j], i)]++
		}

		//整理桶保证稳定
		beforeTimes := 0
		for p := 0; p < len(holder); p++ {
			holder[p] += beforeTimes
			beforeTimes = holder[p]
		}

		//按整理后的桶排序，为了避免在arr内原地交换导致在迭代时错误运算，需要一个中间结果
		res := make([]int, len(arr))
		for q := len(arr) - 1; q >= 0; q-- {
			index := getNumByPosition(arr[q], i)
			//交换位置
			res[holder[index]-1] = arr[q]
			holder[index]--
		}
		arr = res
	}

	return arr
}

func getNumLength(num int) int {
	length := 0
	for ; num > 0; num /= 10 {
		length++
	}
	return length
}

//获取整数对应位置的数值,如果这个位置大于数值长度 返回0
func getNumByPosition(num int, position int) int {
	numLength := getNumLength(num)

	if numLength < position {
		return 0
	} else {
		//-1是因为比如8位长度的数字，第8位其实是在int[7]的位置
		xx, _ := strconv.Atoi(strconv.Itoa(num)[position-1 : position])
		return xx
	}
}
