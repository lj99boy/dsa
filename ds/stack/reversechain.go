package stack

type Chain struct {
	val  int
	next *Chain
}

func ReverseChain(c *Chain) *Chain {
	var holder []*Chain

	for inChain := c; inChain != nil; inChain = inChain.next {
		holder = append(holder, inChain)
	}

	for i := len(holder) - 1; i > 1; i-- {
		holder[i].next = holder[i-1]
	}

	return holder[len(holder)-1]
}