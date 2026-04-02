package main

type MyCircularQueue struct {
	Buff       []int
	Size       int
	FrontIndex int
	RearIndex  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Buff: make([]int, k+1), Size: k + 1, FrontIndex: 0,
		RearIndex: 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Buff[this.RearIndex] = value
	this.RearIndex = (this.RearIndex + 1) % this.Size
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.FrontIndex = (this.FrontIndex + 1) % this.Size
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Buff[this.FrontIndex]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Buff[(this.RearIndex+this.Size-1)%this.Size]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.FrontIndex == this.RearIndex
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.FrontIndex+this.Size-this.RearIndex)%this.Size == 1
}
