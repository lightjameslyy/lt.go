package main

/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 *
 * https://leetcode.com/problems/candy/description/
 *
 * algorithms
 * Hard (27.68%)
 * Total Accepted:    96.6K
 * Total Submissions: 348.8K
 * Testcase Example:  '[1,0,2]'
 *
 * There are N children standing in a line. Each child is assigned a rating
 * value.
 *
 * You are giving candies to these children subjected to the following
 * requirements:
 *
 *
 * Each child must have at least one candy.
 * Children with a higher rating get more candies than their neighbors.
 *
 *
 * What is the minimum candies you must give?
 *
 * Example 1:
 *
 *
 * Input: [1,0,2]
 * Output: 5
 * Explanation: You can allocate to the first, second and third child with 2,
 * 1, 2 candies respectively.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,2]
 * Output: 4
 * Explanation: You can allocate to the first, second and third child with 1,
 * 2, 1 candies respectively.
 * ⁠            The third child gets 1 candy because it satisfies the above two
 * conditions.
 *
 *
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
思路：先从左到右，往左看；然后从右到左，往右看
*/
func candy(ratings []int) int {
	size := len(ratings)
	candy := make([]int, size)
	for i := 0; i < size; i++ {
		candy[i] = 1
	}
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	for i := size - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = max(candy[i], candy[i+1]+1)
		}
	}
	res := 0
	for i := 0; i < size; i++ {
		res += candy[i]
	}
	return res
}
