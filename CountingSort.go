package main

import "slices"

func GetNormValue(value int, minValue int) int {
	return value - minValue
}

func CountingSort(data []int) {
	minData := slices.Min(data)
	counts := make([]int, GetNormValue(slices.Max(data), minData)+1)
	for i := 0; i < len(counts); i++ {
		counts[i] = 0
	}
	for i := 0; i < len(data); i++ {
		counts[GetNormValue(data[i], minData)]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i] + counts[i-1]
	}
	dataCopy := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		dataCopy[counts[GetNormValue(data[i], minData)]-1] = data[i]
		counts[GetNormValue(data[i], minData)]--
	}
	for i := 0; i < len(data); i++ {
		data[i] = dataCopy[i]
	}

}
