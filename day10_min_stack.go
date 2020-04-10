package main

type item struct {
	value int
	next  *item
	min   int
}
type MinStack struct {
	TopItem *item
	Size    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {

	if this.TopItem == nil {
		this.TopItem = &item{value: x, min: x}
	} else {
		min := min(x, this.TopItem.min)
		this.TopItem = &item{value: x, next: this.TopItem, min: min}
	}

	this.Size++
}

func (this *MinStack) Pop() {

	t := this.TopItem.next

	this.TopItem.next = nil
	this.TopItem = t

	this.Size--

}

func (this *MinStack) Top() int {

	return this.TopItem.value
}

func (this *MinStack) GetMin() int {

	if this.TopItem == nil {
		return -1
	}
	return this.TopItem.min
}

func min(a, b int) int {

	if a < b {
		return a
	}

	return b
}

// /**
//  * Your MinStack object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Push(x);
//  * obj.Pop();
//  * param_3 := obj.Top();
//  * param_4 := obj.GetMin();
//  */
