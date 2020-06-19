package sort

// SelectSort an array
func SelectSort(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		minIdx := i
		min := arr[i]
		for j := i + 1; j < size; j++ {
			if arr[j] < min {
				minIdx = j
				min = arr[j]
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
