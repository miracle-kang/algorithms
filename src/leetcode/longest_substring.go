package leetcode

// #3. 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	last := make([]int, 128)
	for i := 0; i < 128; i++ {
		last[i] = -1
	}

	max := 0
	start := 0
	for i, c := range s {
		start = maxNum(last[c]+1, start)
		max = maxNum(i-start+1, max)
		last[c] = i
	}
	return max
}

func maxNum(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
