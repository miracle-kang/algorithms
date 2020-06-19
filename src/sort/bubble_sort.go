package sort

// BubbleSort an array
func BubbleSort(arr []int) {
	size := len(arr)

	for i := 0; i < size; i++ {
		flag := false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
