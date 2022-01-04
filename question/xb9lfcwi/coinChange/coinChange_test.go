package coinChange

import "testing"

func Test_coinChange(t *testing.T) {
	coins := []int{1,2,5}
	amount := 11
	println(coinChange(coins,amount))
}
