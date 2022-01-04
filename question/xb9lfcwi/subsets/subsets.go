package subsets

func subsets(nums []int) [][]int {
	tRes := make([][]int,0)
	recur(&tRes,[]int{},nums,0)

	return tRes
}

func recur(res *[][]int,path []int,choices []int,tI int){
	*res = append(*res,path)

	for i:=tI;i<len(choices);i++ {
		tPath := path
		tPath = append(tPath,choices[i])
		recur(res,tPath,choices,i+1)
	}
}