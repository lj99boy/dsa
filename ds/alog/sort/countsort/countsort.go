package countsort

//计数排序，先准备好一个有序的桶，桶下标就是元素值（优化后是每个元素减去偏移后的值（偏移值就是min）），装完后输出桶就可以了
//一个为了稳定排序的优化，每个桶对应值要加上前面桶对应的值（元素出现次数），然后再循环输入数组找到对应桶，得到自己的次序后，对桶值减一
//循环完输入数组就可以得到结果
//https://yq.aliyun.com/articles/655110
func CountSort(arr []int) []int {
	max := 0
	min := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	holder := make([]int, max-min+1)
	offset := min

	//先放进排序桶里面（排序桶装的是元素出现次数）
	for i := 0; i < len(arr); i++ {
		holder[arr[i]-offset]++
	}

	//优化排序桶结构，每个桶出现次数加上前面的桶出现次数，也就是为了稳定排序结构，因为桶本身是有序的，其实现在装的结果也是对应元素的次序
	beforeTimes := holder[0]
	for i := 1; i < len(holder); i++ {
		if holder[i] != 0 {
			holder[i] += beforeTimes
			beforeTimes = holder[i]
		}
	}

	//输出结果
	//var res []int
	res := make([]int, len(arr))
	//注意这里需要倒序，因为做了--操作后在遇到的相同的数确实是在之前的数的前面的，比如5,4,7,2,7  这里必须倒序才能保证第二次遇到7是在第一次遇到7的前面（--操作已经让第二次操作的遇到相同的数排在前面了）
	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i] - min
		//res = append(res[:holder[index]], append([]int{arr[i]}, res[holder[index]:]...)...)
		//这里的-1是因为比如十个元素的第十位其实是在数组里第九的位置（数组从0开始）
		res[holder[index]-1] = arr[i]
		holder[index]--
	}

	return res
}
