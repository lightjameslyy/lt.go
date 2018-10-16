package leet_0476

import "testing"

// 表格驱动测试
var tests = []struct {
	num int
	res int
}{
	{5, 2},
	{1, 0},
}

func TestFindComplement(t *testing.T) {
	for _, tt := range tests {
		res := FindComplement(tt.num)
		if res != tt.res {
			t.Errorf("input: %d, expected %d, but got %d", tt.num, tt.res, res)
		}
	}
}
