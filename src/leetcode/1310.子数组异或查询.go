/*
 * @lc app=leetcode.cn id=1310 lang=golang
 *
 * [1310] 子数组异或查询
 */

// @lc code=start
func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i := range arr {
		xors[i+1] = xors[i] ^ arr[i]
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		result[i] = xors[query[0]] ^ xors[query[1]+1]
	}
	return result
}

// @lc code=end