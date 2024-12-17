package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	data := make([]int, 0, 10)
	for i := 0; i < cap(data); i++ {
		data = append(data, rand.Intn(20))
	}
	slices.Sort(data)
	fmt.Println(data)
	fmt.Println(BinarySearch(data, 5))
}
