package sort

// BinarySearch a value from arr
// arr: 	The array search for
// value: 	The search value
// left: 	Option search the left or right if there any duplicate value
func BinarySearch(arr []int, value int, left bool) int {
	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)>>1
		if value < arr[mid] {
			high = mid - 1
		} else if value > arr[mid] {
			low = mid + 1
		} else {
			if left {
				if mid == 0 || arr[mid-1] < value {
					return mid
				}
				high = mid - 1
			} else {
				if mid == len(arr)-1 || arr[mid+1] > value {
					return mid
				}
				low = mid + 1
			}
		}
	}
	return -1
}

