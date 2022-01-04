package fib


func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1{
		return 1
	}

	pprev := 0
	prev := 1
	var sum int
	for i:=2;i<=n;i++{
		sum = pprev+prev
		pprev=prev
		prev=sum
	}

	return sum
}

func recur(num int) int {

}