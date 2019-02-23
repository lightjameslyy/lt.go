package main

/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (25.88%)
 * Total Accepted:    79.9K
 * Total Submissions: 308.6K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete at most k
 * transactions.
 *
 * Note:
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [2,4,1], k = 2
 * Output: 2
 * Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit
 * = 4-2 = 2.
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,6,5,0,3], k = 2
 * Output: 7
 * Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit
 * = 6-2 = 4.
 * Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 =
 * 3.
 *
 */

/*
参考：https://blog.csdn.net/qq508618087/article/details/51678245
跟leetcode 123中的方法差不多，将buy和sell的数量扩展到k即可。
思路：当k >= len/2时，问题就退化成了可以交易任意次了，所以只要将任意两天股票差大于０的加起来即可．

当k < len/2时，可以记录k次交易每次买之后和卖以后最大的利润：

１．第ｉ次买操作买下当前股票之后剩下的最大利润为第(i-1)次卖掉股票之后的利润－当前的股票价格．状态转移方程为：

　　　　buy[i] = max(sell[i-1]- curPrice, buy[i]);

２．第ｉ次卖操作卖掉当前股票之后剩下的最大利润为第ｉ次买操作之后剩下的利润＋当前股票价格．状态转移方程为：

　　　　sell[i] = max(buy[i]+curPrice, sell[i]);
*/

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}
	// equivalent to leetcode 122
	if k*2 >= size-1 {
		maxProfit := 0
		for i := 0; i < len(prices)-1; i++ {
			if prices[i] < prices[i+1] {
				maxProfit += prices[i+1] - prices[i]
			}
		}
		return maxProfit
	}
	buy := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		buy[i] = -1 << 31
	}
	sell := make([]int, k+1)

	for _, price := range prices {
		for i := k; i > 0; i-- {
			sell[i] = maxInt(sell[i], buy[i]+price)
			buy[i] = maxInt(buy[i], sell[i-1]-price)
		}
	}
	if sell[k] < 0 {
		return 0
	}
	return sell[k]
}
