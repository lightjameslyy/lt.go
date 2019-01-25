/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (27.91%)
 * Total Accepted:    182.8K
 * Total Submissions: 655K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 *
 * Example 1:
 *
 *
 * Input: [1,2,0]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [3,4,-1,1]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: [7,8,9,11,12]
 * Output: 1
 *
 *
 * Note:
 *
 * Your algorithm should run in O(n) time and uses constant extra space.
 *
 */

/*
 * bucket sort
 * implicit information: the array is expected to be
 *  0  1  2       len-1
 * [1, 2, 3, ..., len]
 * 
 * example: [3,4,-1,1]
 *          
 *          3  4 -1  1
 *          0  1  2  3
 *
 * step 0: i = 0, nums[0] = 3, nums[3-1] == -1 != 3, swap nums[2] and nums[0]
 *
 *          -1  4  3  1
 *          0   1  2  3
 *      
 * step 1: i = 0, nums[0] = -1 < 0
 * 
 *          -1  4  3  1
 *          0   1  2  3
 *
 * step 2: i = 1, nums[1] = 4, nums[4-1] == 1 != 4, swap nums[3] and nums[1]
 *
 *          -1  1  3  4
 *          0   1  2  3
 * 
 * step 3: i = 1, nums[1] = 1, nums[1-1] == 0 != 1, swap nums[1] and nums[0]
 *
 *          1  -1  3  4
 *          0   1  2  3
 *
 * step 4: i = 2, nums[2] = 3, nums[3-1] == 3
 * 
 *          1  -1  3  4
 *          0   1  2  3
 *
 * step 5: i = 3, nums[3] = 4, nums[4-1] == 4
 * 
 *          1  -1  3  4
 *          0   1  2  3
 *
 */
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
