package chain

import "testing"

var chain = &SingleChain{
	val: 1,
	next: &SingleChain{
		val: 2, next: &SingleChain{
			val: 3, next: &SingleChain{
				val: 4, next: &SingleChain{
					val: 5,
				},
			},
		},
	},
}

func TestTraversalSingleChain(t *testing.T) {
	TraversalSingleChain(chain)
}

func TestReverseSingleChain(t *testing.T) {
	ReverseSingleChain(chain)
}
