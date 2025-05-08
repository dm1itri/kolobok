package utils

func CountInSlice[T comparable](slice []T, value T) int {
	count := 0

	for i := range slice {
		if slice[i] == value {
			count++
		}
	}

	return count
}
