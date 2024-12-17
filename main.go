package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var head *Node = nil
	for i := 0; i < 10; i++ {
		head = Insert(head, i)
	}
	head = Delete(head, 8)
	head = Delete(head, 5)
	for i := 0; i < 10; i++ {
		head = Insert(head, rand.Intn(100))
	}
	PreOrder(head)
	fmt.Printf("Height: %d\n", GetMaxHeight(head))
}
