/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (34.59%)
 * Total Accepted:    181.4K
 * Total Submissions: 524.6K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 * 
 * 
 * 
 * 
 * 
 */
func containsNearbyDuplicate(nums []int, k int) bool {
    // m存放数字对应的个数
    m := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if i < k {
            if _, ok := m[nums[i]]; ok {
                return true
            } else {
                m[nums[i]] = true
            }
        } else {
            if _, ok := m[nums[i]]; ok {
                return true
            } else {
                m[nums[i]] = true
                delete(m, nums[i-k])
            }
        }
    }
    return false
}
