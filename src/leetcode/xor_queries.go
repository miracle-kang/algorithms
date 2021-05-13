package leetcode

// #1310. 子数组异或查询
// 有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。
// 对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
// 并返回一个包含给定查询 queries 所有结果的数组。
func xorQueries(arr []int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i, query := range queries {
		xor := arr[query[0]]
		for j := query[0] + 1; j <= query[1]; j++ {
			xor ^= arr[query[j]]
		}
		result[i] = xor
	}
	return result
}

// 优化实现方式
// 官方题解
func xorQueries1(arr []int, queries [][]int) []int {
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
