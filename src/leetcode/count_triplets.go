package leetcode

// 1442. 形成两个异或相等数组的三元组数目
// 给你一个整数数组 arr 。
// 现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length)
// a 和 b 定义如下：
// 		a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
// 		b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
// 请返回能够令 a == b 成立的三元组 (i, j , k) 的数目。
func countTriplets(arr []int) int {
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

// 优化：二重循环
// 当 xors[i] == xors[k]+1 成立时，[i+1, k][i+1,k] 的范围内的任意 jj 都是符合要求的，
// 对应的三元组个数为 k-i。因此我们只需枚举下标 i 和 k。
func countTriplets1(arr []int) int {
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
