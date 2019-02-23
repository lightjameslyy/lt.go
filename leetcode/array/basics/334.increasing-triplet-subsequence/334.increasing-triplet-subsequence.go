package main

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 *
 * https://leetcode.com/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (39.41%)
 * Total Accepted:    84.1K
 * Total Submissions: 213.5K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given an unsorted array return whether an increasing subsequence of length 3
 * exists or not in the array.
 *
 * Formally the function should:
 *
 * Return true if there exists i, j, k
 * such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return
 * false.
 *
 * Note: Your algorithm should run in O(n) time complexity and O(1) space
 * complexity.
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [5,4,3,2,1]
 * Output: false
 *
 *
 *
 */

/*
思路：可以维护一个当前的长度为2的升序序列(小的值叫small, 大的叫big)，
如果碰到比第二个值大的说明可以找到升序的三个值。
并且在过程中不断更新small和big的值，使得他们最小。

NOTE: 在代码中并不用限制i和j的相对位置。
*/
func increasingTriplet(nums []int) bool {
	vi, vj := 1<<31-1, 1<<31-1
	for _, v := range nums {
		if v <= vi {
			vi = v
		} else if v <= vj {
			vj = v
		} else {
			return true
		}
	}
	return false
}
