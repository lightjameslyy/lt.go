/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 *
 * https://leetcode.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (31.04%)
 * Total Accepted:    90.2K
 * Total Submissions: 290.5K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an integer array of size n, find all elements that appear more than ⌊
 * n/3 ⌋ times.
 *
 * Note: The algorithm should run in linear time and in O(1) space.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: [3]
 *
 * Example 2:
 *
 *
 * Input: [1,1,1,3,3,2,2,2]
 * Output: [1,2]
 *
 */
//package main

func majorityElement(nums []int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}
	num1, num2 := 0, 0
	cnt1, cnt2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == num1 {
			cnt1++
		} else if nums[i] == num2 {
			cnt2++
		} else if cnt1 == 0 {
			num1 = nums[i]
			cnt1 = 1
		} else if cnt2 == 0 {
			num2 = nums[i]
			cnt2 = 1
		} else {
			cnt1--
			cnt2--
		}

	}
	cnt1, cnt2 = 0, 0
	for _, val := range nums {
		if val == num1 {
			cnt1++
		} else if val == num2 {
			cnt2++
		}
	}
	if cnt1 > len(nums)/3 {
		res = append(res, num1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, num2)
	}
	return res
}
