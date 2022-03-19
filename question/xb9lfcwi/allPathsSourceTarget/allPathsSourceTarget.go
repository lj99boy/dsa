package allPathsSourceTarget

func allPathsSourceTarget(graph [][]int) [][]int {

	ans = make([][]int,0)

	base := []int{ 0}
	recurse(base,graph,0)

	return ans

}

var ans [][]int

func recurse(path []int,graph [][]int,nowNode int)  {

	if nowNode == len(graph)-1 {
		//tPath := path
		tPath := make([]int,0)
		tPath = append(tPath,path...)
		ans = append(ans,tPath)
		return
	}

	around := graph[nowNode]

	for i:=0;i<len(around);i++{
		path = append(path,around[i])
		recurse(path,graph,around[i])
		path = path[:len(path)-1]
	}

	return
}

