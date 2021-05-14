package leetcode

var romanMap = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9,
	"X": 10, "XL": 40, "L": 50, "XC": 90,
	"C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}

// 13. 罗马数字转整数
func romanToInt(s string) int {
	preVal := 0
	res := 0
	for i := range s {
		val := romanMap[s[i:i+1]]
		if i > 0 && val > preVal {
			res -= preVal
			res += romanMap[s[i-1:i+1]]
		} else {
			res += val
		}
		preVal = val
	}
	return res
}
