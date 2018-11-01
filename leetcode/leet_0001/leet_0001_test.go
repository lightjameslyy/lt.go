package leet_0001

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := TwoSum(nums, 9)
	if len(res) != 2 {
		t.Errorf("2 numbers expected, but got %d numbers", len(res))
	}
	if (res[0] != 0 && res[1] != 1) && (res[0] != 1 && res[1] != 0) {
		t.Errorf("expected: [0, 1] or [1, 0], but got %v", res)
	}
}
