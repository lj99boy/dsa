package dijkstra

import "math"

type nodeState struct {
	id int
	distFromStart int
}

//https://leetcode-cn.com/problems/network-delay-time/
func networkDelayTime(times [][]int, n int, k int) int {

	//建立邻接表
	graph := make(map[int]map[int]int)

	for _,v:=range times{
		if _,ok := graph[v[0]];!ok{
			graph[v[0]] = make(map[int]int)
		}
		graph[v[0]][v[1]] = v[2]
	}

	res := dijkstra(k,graph,n)

	max := 0
	for _,v:=range res{
		if v==math.MaxInt32{
			return -1
		}

		if v>max{
			max = v
		}
	}

	return max
}


// 起点和邻接表
// 邻接表结构为 [目标点][可以到达点1][到达点权重]
//返回起点到所有其他点的最近距离
func dijkstra(start int,graph map[int]map[int]int,n int) []int {
	distTo := make([]int,n+1)
	for i:=1;i<len(distTo);i++{
		distTo[i] = math.MaxInt32
	}

	distTo[start] = 0

	pq:=make([]nodeState,0)

	pq = append(pq,nodeState{id:start,distFromStart: 0})

	for ;len(pq)>0;{
		tItem := pq[0]
		pq = pq[1:]

		if distTo[tItem.id] < tItem.distFromStart {
			continue
		}
		for otherItem,otherItemDist := range graph[tItem.id] {
			tDist:=tItem.distFromStart+otherItemDist
			if tDist<distTo[otherItem]{
				distTo[otherItem] = tDist
				pq = append(pq,nodeState{id:otherItem,distFromStart: tDist})
			}
		}
	}

	return distTo
}
