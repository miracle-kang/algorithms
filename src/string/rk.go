package string

import "math"

// RK 算法的全称叫 Rabin-Karp 算法，是由它的两位发明者 Rabin 和 Karp 的名字来命名的
func RKMatcher(str, matStr string) int {
	n := len(str)
	m := len(matStr)
	matHash := hash(matStr)
	for i := 0; i < n-m; i++ {
		// 对比Hash是否一致
		hash := hash(str[i : i+m])
		if hash != matHash {
			continue
		}
		// Hash 匹配上了以后，再对比一遍防止冲突
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

func hash(str string) int {
	n := len(str)
	res := 0
	for _, c := range str {
		res += int(c) * int(math.Pow10(n-1))
		n--
	}
	return res
}
