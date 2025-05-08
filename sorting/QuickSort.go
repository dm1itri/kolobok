package sorting

func QuickSortLomuto(data []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := data[r]

	i := l - 1

	for j := l; j < r; j++ {
		if data[j] <= pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}

	data[i+1], data[r] = data[r], data[i+1]

	QuickSortLomuto(data, l, i)
	QuickSortLomuto(data, i+2, r)
}

func QuickSortHoare(data []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := data[r] // l || (l + r) / 2 || r || random

	i := l
	j := r

	for i < j {
		for data[i] < pivot {
			i++
		}

		for data[j] > pivot {
			j--
		}

		if i <= j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		}
	}

	QuickSortHoare(data, l, j)
	QuickSortHoare(data, i, r)
}
