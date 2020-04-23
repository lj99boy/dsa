package tree

import "dsa/ds/heap"

func BuildHuffmanTree(weights []int) *ChainBinaryTree {
	priorityQueue := &heap.BinaryHeap{Holder: weights}
	priorityQueue.Build()

	var forest []*ChainBinaryTree
	for ; priorityQueue.Count() > 2; {
		leftNode := priorityQueue.Pop()
		rightNode := priorityQueue.Pop()
		fatherVal := leftNode + rightNode
		priorityQueue.Push(fatherVal)
		tree := &ChainBinaryTree{
			Val:   fatherVal,
			Left:  &ChainBinaryTree{Val: leftNode},
			Right: &ChainBinaryTree{Val: rightNode},
		}

		searched := false
		for i := 0; i < len(forest); i++ {
			if forest[i].Val == tree.Left.Val {
				tree.Left = forest[i]
				forest[i] = tree
				searched = true
			}
			if forest[i].Val == tree.Right.Val {
				tree.Right = forest[i]
				forest[i] = tree
				searched = true
			}
		}

		if searched {
			continue
		} else {
			forest = append(forest, tree)
		}
	}
	return forest[0]
}
