package main

/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 *
 * https://leetcode.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (19.36%)
 * Total Accepted:    85K
 * Total Submissions: 439.1K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * Given an array of integers, find out whether there are two distinct indices
 * i and j in the array such that the absolute difference between nums[i] and
 * nums[j] is at most t and the absolute difference between i and j is at most
 * k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3, t = 0
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1, t = 2
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,5,9,1,5,9], k = 2, t = 3
 * Output: false
 *
 *
 *
 *
 *
 */

// method 1: sort (with indexes)
// time: O(nlogn), space: O(n)

// method 2: TreeSet
// time: O(nlogk), space: O(k)

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// method 3: bucket
// time: O(n), space: O(k)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	size := len(nums)
	if k == 0 || t < 0 || size <= 1 {
		return false
	}
	// t 可以为 0
	// t 需要当除数
	// t++ 后，避免了除0错误。
	// 但 <=t 变成了 <t
	t++
	// 注意，此时已经 t++ 了
	// |a-b|<t
	// -t < a - b < t
	// -1 < a/t - b/t < 1
	// nMap[a/t] = a
	nMap := make(map[int]int, size)
	var ni, m int
	for i := 0; i < size; i++ {
		ni = nums[i]
		m = ni / t
		if ni < 0 {
			// 3 / 4 * 4 == 0, 0 是 3 的左边的 4 的倍数
			// -1/ 4 * 4 == 0, 0 是 -1的右边的 4 的倍数
			// 所以，当除以正数 t 以后，被除数由于正负号的不同，他们的变动方向是不一样的。
			// 当 ni < 0 时，m--，可以统一除法后的变动方向。
			// 可以注释掉这段程序，看看单元测试中，下面 case
			// {
			// 	[]int{-1, 3},
			// 	1,
			// 	3,
			// 	false,
			// },
			m--
		}
		if _, ok := nMap[m]; ok {
			// 说明存在 nums[j]== n, j < i，使得
			// m*t <= ni, n < m*(t+1) 成立
			// 即，|ni-n|<t 成立
			return true
		} else if n, ok := nMap[m-1]; ok && abs(ni, n) < t {
			// 说明存在 nums[j]== n, j < i，使得
			// n < m*t <= ni 成立
			// 这也还是有可能使得 |ni-n|<t 成立的
			// 需要检查 abs(ni, n)
			return true
		} else if n, ok := nMap[m+1]; ok && abs(ni, n) < t {
			// 说明存在 nums[j]== n, j < i，使得
			// ni < m*(t+1) <= n 成立
			// 这也还是有可能使得 |ni-n|<t 成立的
			// 需要检查 abs(ni, n)
			return true
		}

		// nMap[m] = ni
		// 当 t == 4 时 (此 t 已经 t++ 过了)
		// ni == 0,1,2,3 都会使得 m = ni/t = 0
		// 那是否会丢失 ni 的信息呢？导致程序出错呢？
		// 比如，以下case中
		// {
		// 	[]int{3, 0, 6},
		// 	2,
		// 	3, // t == 4
		// 	true,
		// },
		// 3/4 == 0/4 == 0
		// 0 在 3 的后面，且 6-0>4==t
		// 答案是不会
		// 因为程序在遇到 6 之前，就已经因为 3/4 == 0/4 返回 true 了
		nMap[m] = ni
		if i >= k {
			// 删除需不需要检查的点
			delete(nMap, nums[i-k]/t)
		}
	}

	return false
}
