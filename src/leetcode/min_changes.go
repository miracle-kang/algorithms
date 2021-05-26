package leetcode

import "math"

// 1787. 使所有区间的异或结果为零
// 给你一个整数数组 nums​​​ 和一个整数 k​​​​​ 。区间 [left, right]（left <= right）的
// 异或结果 是对下标位于 left 和 right（包括 left 和 right ）之间所有元素进行 XOR 运算的结果
// ：nums[left] XOR nums[left+1] XOR ... XOR nums[right] 。
// 返回数组中 要更改的最小元素数 ，以使所有长度为 k 的区间异或结果等于零。
func minChanges(nums []int, k int) int {
	const maxX = 1 << 10          // x 的范围为 [0, 2^10)
	const inf = math.MaxInt32 / 2 // 极大值，为了防止整数溢出

	n := len(nums)
	f := make([]int, maxX)
	for i := range f {
		f[i] = inf
	}
	// 边界条件 f(-1,0)=0
	f[0] = 0

	for i := 0; i < k; i++ {
		// 第 i 个组的哈希映射
		cnt := map[int]int{}
		size := 0
		for j := i; j < n; j += k {
			cnt[nums[j]]++
			size++
		}

		// 求出 t2
		t2min := arrMin(f...)

		g := make([]int, maxX)
		for j := range g {
			g[j] = t2min
		}
		for mask := range g {
			// t1 则需要枚举 x 才能求出
			for x, cntX := range cnt {
				g[mask] = arrMin(g[mask], f[mask^x]-cntX)
			}
		}

		// 别忘了加上 size
		for j := range g {
			g[j] += size
		}
		f = g
	}

	return f[0]
}

func arrMin(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
