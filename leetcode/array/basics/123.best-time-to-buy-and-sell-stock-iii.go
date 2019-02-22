/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (32.84%)
 * Total Accepted:    137.5K
 * Total Submissions: 418.7K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete at most two
 * transactions.
 *
 * Note: You may not engage in multiple transactions at the same time (i.e.,
 * you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [3,3,5,0,0,3,1,4]
 * Output: 6
 * Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit
 * = 3-0 = 3.
 * Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 =
 * 3.
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you
 * are
 * engaging multiple transactions at the same time. You must sell before buying
 * again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 */

package main

// 找到最大的两个上升段
func maxProfit(prices []int) int {
	maxFirst, maxSecond := 0, 0
	n := len(prices)
	if n < 2 {
		return 0
	}
	// lowIndex和highIndex分别为波谷和波峰的index
	lowIndex, highIndex := -1, -1
	for i := 0; i < n; i++ {
		if i == 0 {
			if prices[0] < prices[1] {
				lowIndex = 0
			}
		} else if i == n-1 {
			if prices[n-1] > prices[n-2] {
				highIndex = n - 1
			}
		} else {
			if prices[i] >= prices[i-1] && prices[i] >= prices[i+1] {
				highIndex = i
			}
			if prices[i] <= prices[i-1] && prices[i] <= prices[i+1] {
				lowIndex = i
			}
		}
		if lowIndex != -1 && highIndex != -1 && lowIndex < highIndex {
			profit := prices[highIndex] - prices[lowIndex]
			if profit > maxFirst {
				maxSecond = maxFirst
				maxFirst = profit
			} else if profit > maxSecond {
				maxSecond = profit
			}
			lowIndex = -1
			highIndex = -1
		}
	}
	return maxFirst + maxSecond
}
