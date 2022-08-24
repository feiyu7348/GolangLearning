// author:zfy  date:2022/8/24

package main

type MyQueue struct {
	stack []int
	back  []int
}

// Constructor /** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

// Push /** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, x)
}

// Pop /** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
		return 0
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

// Peek /** Get the front element. */
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == 0 {
		return 0
	}
	this.back = append(this.back, val)
	return val
}

// Empty /** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}
