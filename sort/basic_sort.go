package sort

func InsertionSort(arr []int) []int {
	if arr != nil && len(arr) > 1 {
		for i := 1; i < len(arr); i++ {
			key := arr[i]
			j := i - 1
			for j >= 0 && key < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = key
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	if arr != nil && len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			minIndex := i
			for j := i + 1; j < len(arr); j++ {
				if arr[j] < arr[minIndex] {
					minIndex = j
				}
			}
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

func BubbleSort(arr []int) []int {
	if arr != nil && len(arr) > 1 {
		swap := true
		for swap {
			swap = false
			for j := 0; j < len(arr)-1; j++ {
				if arr[j+1] < arr[j] {
					arr[j+1], arr[j] = arr[j], arr[j+1]
					swap = true
				}
			}
		}
	}
	return arr
}
