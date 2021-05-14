package leetcode

import (
	"testing"
)

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
