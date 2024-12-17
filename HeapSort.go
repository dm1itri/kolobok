package main

func SiftUp(i int, data []int) {
	for i > 0 && data[i] < data[(i-1)/2] {
		data[i], data[(i-1)/2] = data[(i-1)/2], data[i]
		i = (i - 1) / 2
	}
}

func SiftDown(index int, length int, data []int) {
	for index < length {
		i := index
		l := index*2 + 1
		r := index*2 + 2
		if l < length && data[i] > data[l] {
			i = l
		}
		if r < length && data[i] > data[r] {
			i = r
		}
		if i == index {
			break
		}
		data[index], data[i] = data[i], data[index]
		index = i
	}
}

func HeapSort(data []int) {
	length := len(data)
	for i := length / 2; i >= 0; i-- {
		SiftUp(i, data)
		SiftDown(i, length, data)
	}
	dataCopy := make([]int, length)
	for i := 0; i < length; i++ {
		dataCopy[i] = data[0]
		data[0], data[length-i-1] = data[length-i-1], data[0]
		SiftDown(0, length-i-1, data)
	}
	for i := 0; i < len(dataCopy); i++ {
		data[i] = dataCopy[i]
	}
}
