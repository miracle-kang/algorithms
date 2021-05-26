package leetcode

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	var condition string
	var str string

	condition = "3[a]2[bc]"
	str = decodeString(condition)
	if str != "aaabcbc" {
		t.Error("Condition ", condition, " unexpected result: ", str)
	}

	condition = "3[a2[c]]"
	str = decodeString(condition)
	if str != "accaccacc" {
		t.Error("Condition ", condition, " unexpected result: ", str)
	}

	condition = "2[abc]3[cd]ef"
	str = decodeString(condition)
	if str != "abcabccdcdcdef" {
		t.Error("Condition ", condition, " unexpected result: ", str)
	}

	condition = "abc3[cd]xyz"
	str = decodeString(condition)
	if str != "abccdcdcdxyz" {
		t.Error("Condition ", condition, " unexpected result: ", str)
	}
}

func TestCanCross(t *testing.T) {
	list := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Println(canCross(list[:]))
}

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Error("Unexcept value")
	}
	if romanToInt("IV") != 4 {
		t.Error("Unexcept value")
	}
	if romanToInt("MCMXCIV") != 1994 {
		t.Error("Unexcept value")
	}
}

func TestNumWays(t *testing.T) {
	fmt.Println(numWays(4, 2))
}

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

func TestReverseParentheses(t *testing.T) {
	println(reverseParentheses("(ed(et(oc))el)"))
}
