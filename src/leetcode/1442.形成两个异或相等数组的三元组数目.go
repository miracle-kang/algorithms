/*
 * @lc app=leetcode.cn id=1442 lang=golang
 *
 * [1442] 形成两个异或相等数组的三元组数目
 */

// @lc code=start
func countTriplets(arr []int) int {
	n := len(arr)
	xors := make([]int, n+1)
	for i := 0; i < n; i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}

	res := 0
	for i := 0; i < n-1; i++ {
		for k := i + 1; k < n; k++ {
			if xors[i] == xors[k+1] {
				res += k - i
			}
		}
	}
	return res
}

// @lc code=end

func countTriplets1(arr []int) int {
	n := len(arr)
	xors := make([]int, n+1)
	for i := 0; i < n; i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}

	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if xors[i] == xors[k+1] {
					res++
				}
			}
		}
	}
	return res
}