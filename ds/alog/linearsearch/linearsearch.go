//线性查找
package linearsearch

func LinearSearch(data []int, val int) int {
	for key, v := range data {
		if v == val {
			return key
		}
	}

	return -1
}
