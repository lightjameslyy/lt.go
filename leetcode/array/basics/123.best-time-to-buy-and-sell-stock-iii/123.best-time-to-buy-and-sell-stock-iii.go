package main

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

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// DP: update buy1, sell1, buy2, sell2
func maxProfit(prices []int) int {
	buy1 := -1 << 31 // buy1: the profit after first buy
	sell1 := 0       // sell1: the profit after first sell
	buy2 := -1 << 31 // buy2: the profit after second buy, take sell1 into account
	sell2 := 0       // sell2: the profit after second sell
	// update order: sell2, buy2, sell1, buy1
	for _, price := range prices {
		sell2 = maxInt(sell2, buy2+price)
		buy2 = maxInt(buy2, sell1-price)
		sell1 = maxInt(sell1, buy1+price)
		buy1 = maxInt(buy1, -price)
	}
	// sell2 > buy2+price > (sell1-price)+price > sell1
	if sell2 < 0 {
		return 0
	}
	return sell2
}
