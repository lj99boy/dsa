package coinChange

import "math"

func coinChange(coins []int, amount int) int {
	tDp := make(map[int]int)

	res := runDp(coins,&tDp,amount)
	if res == 998 {
		return -1
	}else{
		return res
	}
}

func runDp (coins []int,dp *map[int]int,amount int) int {
	if amount == 0 {
		return 0
	}
	if amount <0 {
		return -1
	}

	if v,ok:=(*dp)[amount];ok {
		return v
	}

	res := 998

	for i:=0;i<len(coins);i++{
		tRes := runDp(coins,dp,amount-coins[i])
		if tRes == -1 {
			continue
		}
		res = int(math.Min(float64(tRes+1), float64(res)))
	}

	(*dp)[amount] = res
	return res
}