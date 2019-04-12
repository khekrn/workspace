package sort

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	res := make([]int, size, size)
	k := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
			k++
		} else {
			res[k] = right[j]
			j++
			k++
		}
	}

	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}

	return res
}

// MergeSort implementation O(n log n)
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
