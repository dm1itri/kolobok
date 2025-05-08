package sorting

import "slices"

func CountingSort(data []int) {
	rangeValues := slices.Max(data) - slices.Min(data) + 1
	counts := make([]int, 0, rangeValues)

	for i := 0; i < rangeValues; i++ {
		counts[i] = 0
	}

	for i := 0; i < len(data); i++ {
		counts[data[i]]++
	}

	for i := 1; i < rangeValues; i++ {
		counts[i] = counts[i] + counts[i-1]
	}

	dataCopy := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		dataCopy[counts[data[i]]-1] = data[i]
		counts[data[i]]--
	}

	for i := 0; i < len(data); i++ {
		data[i] = dataCopy[i]
	}
}
