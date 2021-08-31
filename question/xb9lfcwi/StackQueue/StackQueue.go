package StackQueue

type MyQueue struct {
	lStack []int
	rStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{lStack: make([]int, 0), rStack: make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if len(this.rStack) > 0 {
		for i:=len(this.rStack)-1;i>=0;i--{
			this.lStack = append(this.lStack,this.rStack[i])
		}
		this.rStack = make([]int, 0)
	}
	this.lStack = append(this.lStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.rStack) > 0 {
		t := this.rStack[len(this.rStack)-1]
		this.rStack = this.rStack[:len(this.rStack)-1]
		return t
	} else{
		this.rStack = make([]int, 0)
		for i:=len(this.lStack)-1;i>=0;i--{
			this.rStack = append(this.rStack,this.lStack[i])
		}
		this.lStack = make([]int, 0)

		t := this.rStack[len(this.rStack)-1]
		this.rStack = this.rStack[:len(this.rStack)-1]
		return t
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.rStack) < 1 {
		this.rStack = make([]int, 0)
		for i:=len(this.lStack)-1;i>=0;i--{
			this.rStack = append(this.rStack,this.lStack[i])
		}
		this.lStack = make([]int, 0)
	}
	return this.rStack[len(this.rStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.lStack) < 1 && len(this.rStack) < 1 {
		return true
	}

	return false
}
