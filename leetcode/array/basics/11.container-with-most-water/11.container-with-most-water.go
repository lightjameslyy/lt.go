package main

// import "fmt"

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (42.40%)
 * Total Accepted:    315.6K
 * Total Submissions: 744.2K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 *
 * Note: You may not slant the container and n is at least 2.
 *
 *
 *
 *
 *
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49.
 *
 *
 *
 * Example:
 *
 *
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 *
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func area(height []int, l, h int) int {
	if l >= h {
		return 0
	}
	return (h - l) * min(height[l], height[h])
}

/*
思路：two pointers
两边比较，小的淘汰
*/
func maxArea(height []int) int {
	size := len(height)
	if size < 2 {
		return 0
	}
	l, h := 0, size-1
	maxArea := 0
	for l < h {
		maxArea = max(maxArea, area(height, l, h))
		if height[l] < height[h] {
			l++
		} else {
			h--
		}
		// fmt.Println(l, h, maxArea)
	}
	return maxArea
}
