package tree

import "dsa/ds/heap"

// 构建哈夫曼树，哈夫曼树就是所有叶子权值乘上它的路径的和最小的树
// 构建过程就是把所有叶子节点权值放入优先队列，每次弹出两个加上和得到父节点值，组成一个树，并把父节点放回优先队列，组成的树去与其他拥有相同节点的树融合，或者单独作为一棵树（一般是在刚开始出现这个情况）
// https://blog.csdn.net/bjweimengshu/article/details/105383513
func BuildHuffmanTree(weights []int) *ChainBinaryTree {
	priorityQueue := &heap.BinaryHeap{Holder: weights}
	priorityQueue.Build()

	var forest []*ChainBinaryTree
	for ; priorityQueue.Count() >= 2; {
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
				break
			}
			if forest[i].Val == tree.Right.Val {
				tree.Right = forest[i]
				forest[i] = tree
				searched = true
				break
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
