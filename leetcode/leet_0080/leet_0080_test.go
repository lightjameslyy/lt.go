package leet_0080

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	res1 := RemoveDuplicates(nums1)
	t.Logf("nums1: %v", nums1)
	if res1 != 5 {
		t.Errorf("expected %d, but got %d", 5, res1)
	}

	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	res2 := RemoveDuplicates(nums2)
	t.Logf("nums2: %v", nums2)
	if res2 != 7 {
		t.Errorf("expected %d, but got %d", 7, res2)
	}

	nums3 := []int{1, 1, 1}
	res3 := RemoveDuplicates(nums3)
	t.Logf("nums3: %v", nums3)
	if res3 != 2 {
		t.Errorf("expected %d, but got %d", 2, res3)
	}
}
