package leetcode

import "testing"

/*
 * TestPalindrome
 *
 *  验证回文串
 */
func TestPalindrome(t *testing.T) {
	var str = "A man, a plan, a canal: Panama"
	var result = isPalindrome(str)

	t.Log(isChar(':'))
	if result != true {
		t.Error("执行结果不符合预期")
	}

	str = "race a car"
	result = isPalindrome(str)
	if result != false {
		t.Error("执行结果不符合预期")
	}
}
