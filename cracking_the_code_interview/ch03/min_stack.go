package main

import "sync"

func main() {

}

type item struct {
	Val      int
	localMin int
}
type MinStack struct {
	lock sync.Mutex
	size int
	s    []item
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.lock.Lock()
	defer this.lock.Unlock()

	localMin := x
	if this.size > 0 && this.s[this.size-1].localMin < x {
		localMin = this.s[this.size-1].localMin
	}
	this.s = append(this.s, item{x, localMin})
	this.size++
}

func (this *MinStack) Pop() {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.size == 0 {
		return
	}
	this.s = this.s[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.s[this.size-1].Val
}

func (this *MinStack) GetMin() int {
	return this.s[this.size-1].localMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
