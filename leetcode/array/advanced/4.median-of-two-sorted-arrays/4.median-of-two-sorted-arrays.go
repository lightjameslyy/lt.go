package main

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (25.50%)
 * Total Accepted:    383.4K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
思路：变形的binary search
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 将复杂度控制在O(log(min(m, n)))
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	// cut1, cut2分别表示nums1和nums2取的左侧元素的个数，范围为[0:len]
	// 在nums1中选定位置cut1后，cut2也确定了，因为：
	//     cut1+cut2=n/2
	// 所以以nums1为基准进行二分即可
	cut1, cut2, l1, r1 := 0, 0, 0, m
	// 最终nums1和nums2中的候选值：
	L1, R1, L2, R2 := 0, 0, 0, 0
	for cut1 <= m {
		cut1 = l1 + (r1-l1)/2
		cut2 = (m+n)/2 - cut1
		// 计算L1，R1，L2，R2，注意边界情况
		// 注意：cut1和cut2是个数，不是下标
		if cut1 == 0 {
			L1 = -1 << 31
		} else {
			L1 = nums1[cut1-1]
		}
		if cut1 == m {
			R1 = 1<<31 - 1
		} else {
			R1 = nums1[cut1]
		}
		if cut2 == 0 {
			L2 = -1 << 31
		} else {
			L2 = nums2[cut2-1]
		}
		if cut2 == n {
			R2 = 1<<31 - 1
		} else {
			R2 = nums2[cut2]
		}
		if L1 > R2 { // cut1太大，左移
			r1 = cut1 - 1
		} else if L2 > R1 { // cut2太大，即cut1太小，右移cut1
			l1 = cut1 + 1
		} else { // 当L1<=R2 && L2<=R1时，满足退出条件
			break
		}
	}
	// 总长度为偶数
	if (m+n)%2 == 0 {
		return float64(max(L1, L2)+min(R1, R2)) / 2
	}

	// 总长度为奇数
	return float64(min(R1, R2))
}
