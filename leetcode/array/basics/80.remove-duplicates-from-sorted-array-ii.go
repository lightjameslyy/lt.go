package main

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
 *
 * algorithms
 * Medium (39.12%)
 * Total Accepted:    181.9K
 * Total Submissions: 465K
 * Testcase Example:  '[1,1,1,2,2,3]'
 *
 * Given a sorted array nums, remove the duplicates in-place such that
 * duplicates appeared at most twice and return the new length.
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * Example 1:
 *
 *
 * Given nums = [1,1,1,2,2,3],
 *
 * Your function should return length = 5, with the first five elements of nums
 * being 1, 1, 2, 2 and 3 respectively.
 *
 * It doesn't matter what you leave beyond the returned length.
 *
 * Example 2:
 *
 *
 * Given nums = [0,0,1,1,1,1,2,3,3],
 *
 * Your function should return length = 7, with the first seven elements of
 * nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.
 *
 * It doesn't matter what values are set beyond the returned length.
 *
 */

/*
 * end: nums[0:end] duplicates at most twice
 * i: traversal pointer
 * count: current duplicates times
 *
 * step 0: end = 2, i = 2, nums[i] == nums[end-2]
 *          e                     e
 *     [1,1,1,2,2,3]   ==>   [1,1,1,2,2,3]
 *          i                     i
 *
 * step 1: end = 2, i = 3, nums[i] != nums[end-2]
 *          e                       e
 *     [1,1,1,2,2,3]   ==>   [1,1,2,2,2,3]
 *            i                     i
 *
 * step 2: end = 3, i = 4, nums[i] != nums[end-2]
 *            e                       e
 *     [1,1,2,2,2,3]   ==>   [1,1,2,2,2,3]
 *              i                     i
 *
 * step 3: end = 4, i = 5, nums[i] != nums[end-2]
 *              e                       e
 *     [1,1,2,2,2,3]   ==>   [1,1,2,2,3,3]
 *                i                     i
 *
 */

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	end := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[end-2] {
			nums[end] = nums[i]
			end++
		}
	}
	return end
}
