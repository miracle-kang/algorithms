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
			high = mid - 1 // 要找的值在左边
		} else if value > arr[mid] {
			low = mid + 1 // 要找的值在右边
		} else {
			if left {
				// 如果有重复元素，查找左边第一个
				if mid == 0 || arr[mid-1] < value {
					return mid
				}
				high = mid - 1
			} else {
				// 如果有重复元素，查找右边第一个
				if mid == len(arr)-1 || arr[mid+1] > value {
					return mid
				}
				low = mid + 1
			}
		}
	}
	return -1
}

// BinarySearchGt 二分查找大于给定值的第一个元素
func BinarySearchGt(arr []int, value int) int {
	size := len(arr)
	low := 0
	high := size - 1
	var mid int

	// 边界情况，第一个元素就大于给定值
	if arr[0] > value {
		return 0
	}

	for low <= high {
		mid = low + (high-low)>>1

		// 要查找的元素在右边，此时中间值小于或等于给定值
		if arr[mid] <= value {
			// 如果中间值右边一位大于给定值，那该元素就是大于给定值的第一个元素
			if mid != size-1 && arr[mid+1] > value {
				return mid + 1
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// BinarySearchLt 二分查找小于给定值的最后一个元素
func BinarySearchLt(arr []int, value int) int {
	size := len(arr)
	low := 0
	high := size - 1
	var mid int

	// 边界情况，最后一个元素都小于给定值
	if arr[size-1] < value {
		return size - 1
	}

	for low <= high {
		mid = low + (high-low)>>1

		// 要查找的元素在左边，此时中间值大于或等于给定值
		if arr[mid] >= value {
			// 如果中间值左边一位小于给定值，那该元素就是小于给定值的最后一个元素
			if mid != 0 && arr[mid-1] < value {
				return mid - 1
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
