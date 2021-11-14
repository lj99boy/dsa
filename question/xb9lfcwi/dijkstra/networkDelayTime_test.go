package dijkstra

import "testing"

func Test_networkDelayTime(t *testing.T) {
	tTime:=make([][]int,0)

	tTime = append(tTime,[]int{2,1,1})
	tTime = append(tTime,[]int{2,3,1})
	tTime = append(tTime,[]int{3,4,1})

	println(networkDelayTime(tTime,4,2))
}
