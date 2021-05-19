package leetcode

import "sort"

// 1738. 找出第 K 大的异或坐标值
// 给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。
// 矩阵中坐标 (a, b) 的 值 可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。
// 请你找出 matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。
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
