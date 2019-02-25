package main

/*
 * @lc app=leetcode id=164 lang=golang
 *
 * [164] Maximum Gap
 *
 * https://leetcode.com/problems/maximum-gap/description/
 *
 * algorithms
 * Hard (31.78%)
 * Total Accepted:    65.9K
 * Total Submissions: 207.3K
 * Testcase Example:  '[3,6,9,1]'
 *
 * Given an unsorted array, find the maximum difference between the successive
 * elements in its sorted form.
 *
 * Return 0 if the array contains less than 2 elements.
 *
 * Example 1:
 *
 *
 * Input: [3,6,9,1]
 * Output: 3
 * Explanation: The sorted form of the array is [1,3,6,9], either
 * (3,6) or (6,9) has the maximum difference 3.
 *
 * Example 2:
 *
 *
 * Input: [10]
 * Output: 0
 * Explanation: The array contains less than 2 elements, therefore return 0.
 *
 * Note:
 *
 *
 * You may assume all elements in the array are non-negative integers and fit
 * in the 32-bit signed integer range.
 * Try to solve it in linear time/space.
 *
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

func getMinMax(nums []int) (minVal, maxVal int) {
	minVal, maxVal = 1<<31-1, -1<<31
	for _, v := range nums {
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}
	return
}

func fill(nums []int, val int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = val
	}
}

/*
思路：bucket
*/
func maximumGap(nums []int) int {
	size := len(nums)
	if size < 2 {
		return 0
	}
	minVal, maxVal := getMinMax(nums)
	bktWidth := (maxVal-minVal)/size + 1
	bktMin, bktMax := make([]int, size), make([]int, size)
	fill(bktMin, 1<<31-1)
	fill(bktMax, -1<<31)
	for _, v := range nums {
		i := (v - minVal + 1) / bktWidth
		bktMin[i] = min(bktMin[i], v)
		bktMax[i] = max(bktMax[i], v)
	}
	res := -1 << 31
	preMax := bktMax[0]
	for i := 1; i < size; i++ {
		if bktMax[i] != 1<<31-1 {
			if preMax != -1<<31 {
				res = max(res, bktMin[i]-preMax)
			}
			preMax = bktMax[i]
		}
	}
	return res
}
