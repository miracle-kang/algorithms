package string

import "testing"

func TestBFMatcher(t *testing.T) {
	if BFMatcher("abcabcaaaa", "cabc") != 2 {
		t.Error("BFMatcher failed.")
	}
	if BFMatcher("abcabcaaaa", "bbbb") != -1 {
		t.Error("BFMatcher failed.")
	}
}

func TestRKMatcher(t *testing.T) {
	if RKMatcher("abcabcaaaa", "cabc") != 2 {
		t.Error("BFMatcher failed.")
	}
	if RKMatcher("abcabcaaaa", "bbbb") != -1 {
		t.Error("BFMatcher failed.")
	}
}

func TestBMMatcher(t *testing.T) {

}

func TestKMPMatcher(t *testing.T) {

}
