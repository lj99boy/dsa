package tu

var pathHolder [][]int

func allPathsSourceTarget(graph [][]int) [][]int {

	pathHolder := make([][]int, 0)

	ttPath:= make([]int,0)

	rFunc(0,graph[0],ttPath,graph)

	return pathHolder
}

func rFunc(key int, iGraph []int, tPath []int, graph [][]int) {
	if len(iGraph) < 1 {
		tPath = append(tPath, key)
		x := make([]int,0)
		copy(x,tPath)
		pathHolder = append(pathHolder, x)
		return
	}

	tPath = append(tPath, key)
	for k, v := range iGraph {
		rFunc(k, graph[v], tPath, graph)
	}

	tPath = tPath[0 : len(tPath)-1]
}
