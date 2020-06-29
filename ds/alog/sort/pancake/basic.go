package pancake

func Basic(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		maxI := 0
		for j := i; j > 0; j-- {
			if arr[j] > arr[maxI] {
				maxI = j
			}
		}
		flip(arr[0 : maxI+1])
		flip(arr[0 : i+1])
	}
	return arr
}

//切片内每位反转，如第一位和最后一位，第二位和倒数第二位
func flip(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
