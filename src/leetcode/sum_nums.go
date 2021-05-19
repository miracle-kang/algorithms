package leetcode

// 剑指 Offer #64
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
func sumNums(n int) int {
	var sum func(int) bool
	res := 0
	sum = func(m int) bool {
		res += m
		return m > 0 && sum(m-1)
	}
	sum(n)
	return res
}
