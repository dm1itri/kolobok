package main

func BinarySearch(data []int, key int) int {
	l := 0
	r := len(data) - 1
	for l <= r {
		mid := (r + l) / 2
		if data[mid] < key {
			l = mid + 1
		} else if data[mid] > key {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
