package sort

// MergeSort an array
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)>>1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	l, r, k := left, mid+1, 0

	for i := left; i <= right; i++ {
		if l > mid {
			temp[k] = arr[r]
		} else if r > right {
			temp[k] = arr[l]
			l++
		} else {
			if arr[l] <= arr[r] {
				temp[k] = arr[l]
				l++
			} else {
				temp[k] = arr[r]
				r++
			}
		}
		k++
	}

	for i := 0; i < right-left+1; i++ {
		arr[left+i] = temp[i]
	}
}
