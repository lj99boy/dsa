package main

var pathHolder [][]int

func allPathsSourceTarget(graph [][]int) [][]int {

	pathHolder = make([][]int, 0)

	ttPath := make([]int, 0)

	rFunc(0, graph[0], ttPath, graph)

	//fmt.Printf("%v",pathHolder)
	return pathHolder
}

func rFunc(key int, iGraph []int, tPath []int, graph [][]int) {
	tPath = append(tPath, key)

	if key == len(graph)-1 {
		//x := make([]int, 0)
		//copy(x, tPath)
		xx := make([]int,0)
		xx = append(xx,tPath...)
		pathHolder = append(pathHolder, xx)
		tPath = tPath[0 : len(tPath)-1]
		return
	}

	//tPath = append(tPath, key)
	for _, v := range iGraph {
		rFunc(v, graph[v], tPath, graph)
	}

	tPath = tPath[0 : len(tPath)-1]
}

func main() {
	x := make([][]int, 0)

	x = append(x, []int{1, 2})
	x = append(x, []int{3})
	x = append(x, []int{3})
	x = append(x, []int{})

	allPathsSourceTarget(x)

}
