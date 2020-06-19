package sort

// InsertSort an array
func InsertSort(arr []int) {
	size := len(arr)
	for i := 1; i < size; i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
}
