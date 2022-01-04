package hitBallon

import "testing"

func Test_hitBallon(t *testing.T) {
	tSamples := make([][]int,0)

	tSamples = append(tSamples,[]int{4,8})
	tSamples = append(tSamples,[]int{2,6})
	tSamples = append(tSamples,[]int{10,18})
	tSamples = append(tSamples,[]int{9,11})

	println(hitBallon(tSamples))
}
