package permute


func Permute(nums []int) [][]int {

	res := [][]int{}

	tPath := make([]int,0)

	var cloFunc func([]int,[]int)
	cloFunc = func(choice []int, path []int) {
		if len(choice) == 0 {
			res = append(res, path)
		}
		for i := 0; i < len(choice); i++ {
			var tmpChoice []int
			var tmpPath = path
			tmpPath = append(tmpPath,choice[i])
			for j:=0;j<i;j++{
				tmpChoice = append(tmpChoice,choice[j])
			}
			for j:=i+1;j<len(choice);j++{
				tmpChoice = append(tmpChoice,choice[j])
			}

			cloFunc(tmpChoice,tmpPath)
			tmpPath = path[:len(tmpPath)-1]
		}

	}

	cloFunc(nums,tPath)

	return res
}

