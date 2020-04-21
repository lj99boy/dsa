package heap

//BinaryHeap 暂定为最小堆
//https://mp.weixin.qq.com/s?__biz=MzIxMjE5MTE1Nw==&mid=2653195207&idx=2&sn=12689c6c1a92e7ec3cce4d423019ec2a&chksm=8c99f91dbbee700b8e760d06b27582037ab0713295dacf2b5a7a7f954c0032fe860aa0bf8b74&mpshare=1&scene=1&srcid=0918SEVlQv0L7wjEETZ9IbS4#rd
type BinaryHeap struct {
	holder []int
}

//Build 使无序二叉堆变有序
func (bh *BinaryHeap) Build() {

}

//时间复杂度n
func (bh *BinaryHeap) Push(val int) {
	bh.holder = append(bh.holder, val)
	for i := len(bh.holder) - 1; i >= 1; {
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
		if bh.holder[i] < bh.holder[fatherIndex] {
			bh.holder[i], bh.holder[fatherIndex] = bh.holder[fatherIndex], bh.holder[i]
			i = fatherIndex
		} else {
			return
		}
	}
}

func (bh *BinaryHeap) Pop() (top int) {
	top = bh.holder[0]
	bh.holder[0] = bh.holder[len(bh.holder)-1]
	bh.holder = bh.holder[0 : len(bh.holder)-1]

	for i := 0; i <= (len(bh.holder)-1)/2; {
		if len(bh.holder)-1 >= i*2+1 && bh.holder[i] > bh.holder[i*2+1] {
			bh.holder[i], bh.holder[i*2+1] = bh.holder[i*2+1], bh.holder[i]
			i = i*2 + 1
		} else if len(bh.holder)-1 >= i*2+2 && bh.holder[i] > bh.holder[i*2+2] {
			bh.holder[i], bh.holder[i*2+2] = bh.holder[i*2+2], bh.holder[i]
			i = i*2 + 2
		} else {
			return
		}
	}
	return
}
