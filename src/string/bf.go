package string

// BF 算法中的 BF 是 Brute Force 的缩写，中文叫作暴力匹配算法，也叫朴素匹配算法。
func BFMatcher(str, matStr string) int {
	n := len(str)
	m := len(matStr)
	for i := 0; i < n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if str[i+j] != matStr[j] {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return -1
}
