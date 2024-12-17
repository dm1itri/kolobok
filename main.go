package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	data := make([]int, 0, 1000000)
	goodData := make([]int, cap(data))
	for i := 0; i < cap(data); i++ {
		data = append(data, rand.Int())
		goodData[i] = data[i]
	}
	QuickSortLomuto(data, 0, len(data)-1)
	//QuickSortHoare(data, 0, len(data)-1)
	slices.Sort(goodData)
	fmt.Println(slices.Equal(data, goodData))
}
