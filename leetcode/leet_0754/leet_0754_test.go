package leet_0754

import "testing"

// 表格驱动测试
var tests = []struct {
	target      int
	reachNumber int
}{
	{3, 2},
	{2, 3},
}

func TestReachNumber(t *testing.T) {
	for _, tt := range tests {
		res := ReachNumber(tt.target)
		if res != tt.reachNumber {
			t.Errorf("input: %d, expected %d, but got %d", tt.target, tt.reachNumber, res)
		}
	}
}
