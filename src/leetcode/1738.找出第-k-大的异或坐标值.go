/*
 * @lc app=leetcode.cn id=1738 lang=golang
 *
 * [1738] 找出第 K 大的异或坐标值
 */

// @lc code=start
func kthLargestValue(matrix [][]int, k int) int {
	n := len(matrix)
	m := len(matrix[0])
	arr := make([]int, 0, m*n)
	xors := make([][]int, n+1)
	xors[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		xors[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			xors[i][j] = xors[i-1][j] ^ xors[i][j-1] ^ xors[i-1][j-1] ^ matrix[i-1][j-1]
			arr = append(arr, xors[i][j])
		}
	}
	return query(arr, k)
}

func query(arr []int, k int) int {
	sort.Ints(arr)
	return arr[len(arr)-k]
}

// @lc code=end

