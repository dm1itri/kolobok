package sorting

func InsertionSort(data []int) {
	for j := 1; j < len(data); j++ {
		key := data[j]
		i := j - 1

		for i >= 0 && key < data[i] {
			data[i+1] = data[i]
			i--
		}

		data[i+1] = key
	}
}
