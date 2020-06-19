package leetcode

import "testing"

func TestDecodeString(t *testing.T) {
	var condition string
	var str string

	condition = "3[a]2[bc]"
	str = decodeString(condition)
	if str != "aaabcbc" {
		t.Error("Condition ", condition, " unexcepted result: ", str)
	}

	condition = "3[a2[c]]"
	str = decodeString(condition)
	if str != "accaccacc" {
		t.Error("Condition ", condition, " unexcepted result: ", str)
	}

	condition = "2[abc]3[cd]ef"
	str = decodeString(condition)
	if str != "abcabccdcdcdef" {
		t.Error("Condition ", condition, " unexcepted result: ", str)
	}

	condition = "abc3[cd]xyz"
	str = decodeString(condition)
	if str != "abccdcdcdxyz" {
		t.Error("Condition ", condition, " unexcepted result: ", str)
	}
}
