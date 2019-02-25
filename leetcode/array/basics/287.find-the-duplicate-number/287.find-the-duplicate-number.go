package main

/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 *
 * https://leetcode.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (48.23%)
 * Total Accepted:    164.9K
 * Total Submissions: 341.9K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * Given an array nums containing n + 1 integers where each integer is between
 * 1 and n (inclusive), prove that at least one duplicate number must exist.
 * Assume that there is only one duplicate number, find the duplicate one.
 *
 * Example 1:
 *
 *
 * Input: [1,3,4,2,2]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [3,1,3,4,2]
 * Output: 3
 *
 * Note:
 *
 *
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n2).
 * There is only one duplicate number in the array, but it could be repeated
 * more than once.
 *
 *
 */

/*
思路：fast-slow pointers
跟有环链表类比
1. 找相遇点
2. 一个从开始出发，一个从相遇点出发，最终相遇处是重复处
*/
func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}
	fast = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
