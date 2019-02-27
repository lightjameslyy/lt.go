package main

/*
 * @lc app=leetcode id=327 lang=golang
 *
 * [327] Count of Range Sum
 *
 * https://leetcode.com/problems/count-of-range-sum/description/
 *
 * algorithms
 * Hard (31.99%)
 * Total Accepted:    29.3K
 * Total Submissions: 91.8K
 * Testcase Example:  '[-2,5,-1]\n-2\n2'
 *
 * Given an integer array nums, return the number of range sums that lie in
 * [lower, upper] inclusive.
 * Range sum S(i, j) is defined as the sum of the elements in nums between
 * indices i and j (i ≤ j), inclusive.
 *
 * Note:
 * A naive algorithm of O(n2) is trivial. You MUST do better than that.
 *
 * Example:
 *
 *
 * Input: nums = [-2,5,-1], lower = -2, upper = 2,
 * Output: 3
 * Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective
 * sums are: -2, -1, 2.
 *
 */

/*
思路1: 暴力, TLE
func countRangeSum(nums []int, lower int, upper int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				res++
			}
		}
	}
	return res
}
*/

/*
思路2: merge sort的变形
*/
func countRangeSum(nums []int, lower int, upper int) int {
	// cache all sums
	sums := make([]int64, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + (int64)(nums[i])
	}
	return mergeSort(sums, 0, len(sums), lower, upper)
}

func mergeSort(sums []int64, start, end, lower, upper int) int {
	if end-start <= 1 {
		return 0
	}
	mid := start + (end-start)/2
	count := mergeSort(sums, start, mid, lower, upper) + mergeSort(sums, mid, end, lower, upper)
	j, k, t := mid, mid, mid
	cache := make([]int64, end-start)
	for i, r := start, 0; i < mid; {
		// 在[mid, end]找到令i到k的和>=lower的最左的位置
		for k < end && sums[k]-sums[i] < int64(lower) {
			k++
		}
		// 在[mid, end]找到令i到j的和>upper的最左的位置
		for j < end && sums[j]-sums[i] <= int64(upper) {
			j++
		}
		count += j - k
		// merge
		for t < end && sums[t] < sums[i] {
			cache[r] = sums[t]
			r++
			t++
		}
		cache[r] = sums[i]
		i++
		r++
	}
	copy(sums[start:t], cache[:t-start])
	return count
}
