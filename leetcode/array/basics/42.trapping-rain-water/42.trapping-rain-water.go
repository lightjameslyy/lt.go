package main

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (41.74%)
 * Total Accepted:    254.2K
 * Total Submissions: 608.9K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 *
 *
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 *
 * Example:
 *
 *
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 *
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
思路：leetcode 11的扩展，依然用two pointers
*/
func trap(height []int) int {
	leftMax, rightMax := 0, 0
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		if height[l] < height[r] { // 隐含着leftMax < rightMax，此时l向右移动
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
			l++
		} else {
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
			r--
		}
	}
	return res
}
