package leet_0004

import (
	"math"
	"testing"
)

// 表格驱动测试
var tests = []struct {
	nums1 []int
	nums2 []int
	res   float64
}{
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, tt := range tests {
		res := FindMedianSortedArrays(tt.nums1, tt.nums2)
		// 如何判断结果是否正确呢？
		// 在这个问题中，最终的结果只可能是：
		// 1. int
		// 2. 两个int的平均值（***.5）
		// 因此，如果结果不对，差距至少是0.4以上（约等于0.5），这里我们可以
		// 用0.1来判断
		if math.Abs(res-tt.res) >= 0.1 {
			t.Errorf("input: nums1:%v, nums2:%v, expected %f, but got %f",
				tt.nums1, tt.nums2, tt.res, res)
		}
	}
}
