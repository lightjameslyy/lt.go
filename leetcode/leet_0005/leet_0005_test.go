package leet_0005

import (
	"testing"
)

// 表格驱动测试
var tests = []struct {
	s   string
	res string
}{
	{"", ""},
	{"a", "a"},
	{"babad", "bab"},
	{"cbbd", "bb"},
}

func TestLongestPalindrome1(t *testing.T) {
	for _, tt := range tests {
		res := LongestPalindrome1(tt.s)
		if tt.res != res {
			t.Errorf("input: s:\"%s\", expected \"%s\", but got \"%s\"",
				tt.s, tt.res, res)
		}
	}
}

func TestLongestPalindrome2(t *testing.T) {
	for _, tt := range tests {
		res := LongestPalindrome2(tt.s)
		if tt.res != res {
			t.Errorf("input: s:\"%s\", expected \"%s\", but got \"%s\"",
				tt.s, tt.res, res)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	for _, tt := range tests {
		res := LongestPalindrome(tt.s)
		if tt.res != res {
			t.Errorf("input: s:\"%s\", expected \"%s\", but got \"%s\"",
				tt.s, tt.res, res)
		}
	}
}
