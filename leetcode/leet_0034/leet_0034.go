package leet_0034

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if nums[mid] != target {
		return []int{-1, -1}
	}

	// find left boundary
	l, r, m := left, mid, 0
	for l <= r {
		if nums[l] == target {
			break
		}
		m = l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	left = l

	// find right boundary
	l, r = mid, right
	for l <= r {
		if nums[r] == target {
			break
		}
		m = l + (r-l+1)/2
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	right = r
	return []int{left, right}
}
