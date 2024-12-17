package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	data := make([]int, 0, 10000000)
	goodData := make([]int, cap(data))
	for i := 0; i < cap(data); i++ {
		data = append(data, rand.Intn(100000))
		goodData[i] = data[i]
	}
	HeapSort(data)
	//fmt.Println(data)
	slices.Sort(goodData)
	//fmt.Println(goodData)
	fmt.Println(slices.Equal(data, goodData))
}
