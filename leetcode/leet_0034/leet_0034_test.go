package leet_0034

import "testing"

// 表格驱动测试
var tests = []struct {
	nums   []int
	target int
	res    []int
}{
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{2, 2}, 2, []int{0, 1}},
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSearchRange(t *testing.T) {
	for _, tt := range tests {
		res := SearchRange(tt.nums, tt.target)
		if equal(res, tt.res) == false {
			t.Errorf("input: %d, expected %d, but got %d", tt.nums, tt.res, res)
		}
	}
}
