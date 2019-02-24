package main

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (40.79%)
 * Total Accepted:    189.4K
 * Total Submissions: 464.4K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers, find the length of the longest
 * consecutive elements sequence.
 *
 * Your algorithm should run in O(n) complexity.
 *
 * Example:
 *
 *
 * Input: [100, 4, 200, 1, 3, 2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 */

/*
思路：hashtable
*/
func longestConsecutive(nums []int) int {
	// 1. 先把所有数组放入set中
	hashSet := make(map[int]bool)
	for _, num := range nums {
		hashSet[num] = true
	}
	res := 0
	// 2. 遍历nums，找到边界，更新res
	for _, num := range nums {
		// num-1不在set中 ==> num是左边界
		if _, ok := hashSet[num-1]; !ok {
			tmp := 1
			// 找右边界
			for i := num + 1; ; i++ {
				if _, ok := hashSet[i]; !ok {
					break
				}
				tmp++
			}
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}
