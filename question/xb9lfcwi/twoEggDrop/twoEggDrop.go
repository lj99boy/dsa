package twoEggDrop

import "math"

var cache map[int]map[int]int

func twoEggDrop(n int) int {
	cache = make(map[int]map[int]int)
	for i:=0;i<=2;i++{
		cache[i] = make(map[int]int)
	}
	return dp(2,n)
}

func dp(eggNum int,floor int) int {
	if floor == 0{
		return 0
	}

	if eggNum == 1 {
		return floor
	}

	if v,ok := cache[eggNum][floor];ok{
		return v
	}

	res := math.MaxInt32
	for i:=1;i<=floor;i++{
		res = int(math.Min(math.Max(float64(dp(eggNum-1, i-1)), float64(dp(eggNum, floor-i))) + 1, float64(res)))
	}
	cache[eggNum][floor] = res

	return res
}