package permute

func permute(nums []int) [][]int {
	res := make([][]int,0)

	recur(&res,[]int{},nums)
	return res
}


func recur(res *[][]int,path []int,choices []int){
	if len(choices) < 1 {
		*res = append(*res,path)
		return
	}
	for i:=0;i<len(choices);i++{
		tPath := path
		tPath = append(tPath,choices[i])
		var tmpChoice []int
		for j:=0;j<i;j++{
			tmpChoice = append(tmpChoice,choices[j])
		}
		for j:=i+1;j<len(choices);j++{
			tmpChoice = append(tmpChoice,choices[j])
		}
		recur(res,tPath,tmpChoice)
	}
}