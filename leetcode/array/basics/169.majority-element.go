/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (51.06%)
 * Total Accepted:    337.2K
 * Total Submissions: 660.4K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */
// package main

/*
 * using hash map
func majorityElement(nums []int) int {
	mp := make(map[int]int)
	for _, val := range nums {
		if count, ok := mp[val]; ok {
			mp[val] = count + 1
		} else {
			mp[val] = 1
		}
		if mp[val] > len(nums)/2 {
			return val
		}
	}
	return nums[0]
}
*/

// one pass
func majorityElement(nums []int) int {
	index, count := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[index] == nums[i] {
			count = count + 1
		} else {
			count = count - 1
		}
		if count == 0 {
			index = i
			count = 1
		}
	}
	return nums[index]
}
