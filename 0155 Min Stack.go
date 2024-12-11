package main

import (
	"fmt"
)

type MinStack struct {
	values    []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	if len(this.minValues) != 0 {
		this.minValues = append(this.minValues, min(this.minValues[len(this.minValues)-1], val))
	} else {
		this.minValues = append(this.minValues, val)
	}

}

func (this *MinStack) Pop() {
	this.values = append(this.values[:len(this.values)-1])
	this.minValues = append(this.minValues[:len(this.minValues)-1])
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValues[len(this.minValues)-1]
}

func main() {
	obj := Constructor()
	obj.Push(10)
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3)
	fmt.Println(param_4)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
