package sorting

func Merge(data []int, l int, m int, r int) {
	i := 0
	j := 0
	dataCopy := make([]int, r-l)

	for l+i < m && m+j < r {
		if data[l+i] <= data[m+j] {
			dataCopy[i+j] = data[l+i]
			i++
		} else {
			dataCopy[i+j] = data[m+j]
			j++
		}
	}

	for l+i < m {
		dataCopy[i+j] = data[l+i]
		i++
	}

	for m+j < r {
		dataCopy[i+j] = data[m+j]
		j++
	}

	for i = 0; i < r-l; i++ {
		data[l+i] = dataCopy[i]
	}
}
func MergeSort(data []int, l int, r int) {
	if l+1 >= r {
		return
	}

	mid := (r + l) / 2
	MergeSort(data, l, mid)
	MergeSort(data, mid, r)
	Merge(data, l, mid, r)
}
