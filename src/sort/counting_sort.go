package sort

// CountingSort an array
func CountingSort(arr []int) {
	size := len(arr)
	max := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 计数
	var c = make([]int, max+1)
	for i := 0; i < size; i++ {
		c[arr[i]]++
	}

	// 累加
	for i := 1; i < len(c); i++ {
		c[i] = c[i-1] + c[i]
	}

	// 排序
	var index int
	var temp = make([]int, size)
	for i := size - 1; i >= 0; i-- {
		index = c[arr[i]] - 1 // 从计数累加数组中取出索引
		temp[index] = arr[i]
		// temp[c[arr[i]-1]] = arr[i]
		c[arr[i]]--
	}

	// 拷贝回原数组
	for i := 0; i < size; i++ {
		arr[i] = temp[i]
	}
}
