package MinStack

type MinStack struct {
	h      []int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{h: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.min) < 1 {
		this.min = append(this.min, x)
	} else {
		if x <= this.min[len(this.min)-1] {
			this.min = append(this.min, x)
		}
	}

	this.h = append(this.h, x)
}

func (this *MinStack) Pop() {
	if this.h[len(this.h)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.h = this.h[:len(this.h)-1]
}

func (this *MinStack) Top() int {
	return this.h[len(this.h)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
