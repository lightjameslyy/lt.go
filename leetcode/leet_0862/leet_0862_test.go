package leet_0862

import "testing"

// 表格驱动测试
var tests = []struct {
	A   []int
	K   int
	res int
}{
	{[]int{1}, 1, 1},
	{[]int{1, 2}, 4, -1},
	{[]int{2, -1, 2}, 3, 3},
}

func TestShortestSubarray(t *testing.T) {
	for _, tt := range tests {
		res := ShortestSubarray(tt.A, tt.K)
		if res != tt.res {
			t.Errorf("input: A = %v, K = %d, expected %d, but got %d", tt.A, tt.K, tt.res, res)
		}
	}
}
