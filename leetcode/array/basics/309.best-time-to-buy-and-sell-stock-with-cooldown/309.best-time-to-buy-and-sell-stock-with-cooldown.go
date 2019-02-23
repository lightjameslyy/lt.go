package main

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (43.41%)
 * Total Accepted:    82.3K
 * Total Submissions: 189.7K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (ie, buy one and sell one share of the stock
 * multiple times) with the following restrictions:
 *
 *
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 * After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1
 * day)
 *
 *
 * Example:
 *
 *
 * Input: [1,2,3,0,2]
 * Output: 3
 * Explanation: transactions = [buy, sell, cooldown, buy, sell]
 *
 */

/*
参考：http://zxi.mytechroad.com/blog/dynamic-programming/leetcode-309-best-time-to-buy-and-sell-stock-with-cooldown/
思路：状态机 + DP
    共有3种状态：hold, sold, rest
    hold -> hold, sold
    sold -> rest
    rest -> rest, hold
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	hold, sold, rest := -1<<31, 0, 0
	for _, price := range prices {
		preSold := sold
		sold = hold + price
		hold = max(hold, rest-price)
		rest = max(rest, preSold)
	}
	return max(rest, sold)
}
