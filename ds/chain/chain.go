package chain

import "fmt"

type SingleChain struct {
	val  interface{}
	next *SingleChain
}

func ReverseSingleChain(chain *SingleChain) {
	TraversalSingleChain(chain)
	var stack []*SingleChain
	for ; chain != nil; {
		stack = append(stack, chain)
		chain = chain.next
	}
	for i := len(stack) - 1; i >= 1; i-- {
		stack[i].next = stack[i-1]
	}
	stack[0].next = nil
	chain = stack[len(stack)-1]
	TraversalSingleChain(chain)

}

func TraversalSingleChain(chain *SingleChain) {
	for ; chain != nil; {
		fmt.Printf("%v-", chain.val)
		chain = chain.next
	}
	println(" ")
}
