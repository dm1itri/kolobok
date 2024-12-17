package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	data := make([]int, 0, 100000000)
	goodData := make([]int, cap(data))
	for i := 0; i < cap(data); i++ {
		data = append(data, rand.Int())
		goodData[i] = data[i]
	}
	MergeSort(data, 0, len(data))
	slices.Sort(goodData)
	fmt.Println(slices.Equal(data, goodData))
}
