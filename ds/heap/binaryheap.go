package heap

//BinaryHeap 暂定为最小堆
//https://mp.weixin.qq.com/s?__biz=MzIxMjE5MTE1Nw==&mid=2653195207&idx=2&sn=12689c6c1a92e7ec3cce4d423019ec2a&chksm=8c99f91dbbee700b8e760d06b27582037ab0713295dacf2b5a7a7f954c0032fe860aa0bf8b74&mpshare=1&scene=1&srcid=0918SEVlQv0L7wjEETZ9IbS4#rd
type BinaryHeap struct {
	//因为是一个完全二叉树，所以可以用数组方便的表示（如果不是完全二叉树，那很多为了满足子节点等于父节点*2+1 父节点*2+2这个条件，要补充很多空位构造完全二叉树
	Holder []int
}

//Build 使无序二叉堆变有序
func (bh *BinaryHeap) Build() {
	for i := len(bh.Holder) / 2; i >= 0; {
		bh.AdjustDown(i)
		i--
	}
}

func (bh *BinaryHeap) Count() int {
	return len(bh.Holder)
}

func (bh *BinaryHeap) IsEmpty() bool {
	return len(bh.Holder) == 0
}

func (bh *BinaryHeap) Push(val int) {
	bh.Holder = append(bh.Holder, val)
	bh.AdjustUp(len(bh.Holder) - 1)
}

//时间复杂度n
func (bh *BinaryHeap) AdjustUp(index int) {
	if index > len(bh.Holder)-1 {
		panic("index should not over length")
	}
	swapTmp := bh.Holder[index]
	//for i := len(bh.Holder) - 1; i >= 1; {
	for i := index; i >= 1; {
		isOdd := false
		if i%2 != 0 {
			isOdd = true
		}
		fatherIndex := 0
		if isOdd {
			fatherIndex = i / 2
		} else {
			fatherIndex = i/2 - 1
		}
		if swapTmp < bh.Holder[fatherIndex] {
			bh.Holder[i] = bh.Holder[fatherIndex]
			i = fatherIndex
		} else {
			//减少交换操作
			bh.Holder[i] = swapTmp
			return
		}
	}
	bh.Holder[0] = swapTmp
}

func (bh *BinaryHeap) AdjustDown(index int) {
	swapTmp := bh.Holder[index]
	for i := index; i <= len(bh.Holder)-1; {
		if i*2+2 > len(bh.Holder)-1 {
			bh.Holder[i] = swapTmp
			return
		}

		minIndex := 0
		if bh.Holder[i*2+1] < bh.Holder[i*2+2] {
			minIndex = i*2 + 1
		} else {
			minIndex = i*2 + 2
		}

		if swapTmp > bh.Holder[minIndex] {
			bh.Holder[i] = bh.Holder[minIndex]
			i = minIndex
		} else {
			bh.Holder[i] = swapTmp
			return
		}
	}
	return
}

func (bh *BinaryHeap) Pop() (top int) {
	top = bh.Holder[0]
	if len(bh.Holder) == 1 {
		bh.Holder = bh.Holder[:len(bh.Holder)-1]
		return
	}
	bh.Holder[0] = bh.Holder[len(bh.Holder)-1]
	bh.Holder = bh.Holder[:len(bh.Holder)-1]
	bh.AdjustDown(0)
	return
}
