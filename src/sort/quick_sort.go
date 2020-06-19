package sort

// QuickSort an array
// arr: The array to sort
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := partition(arr, left, right)
	quickSort(arr, left, mid-1)
	quickSort(arr, mid+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[right], arr[i] = arr[i], pivot
	return i
}
