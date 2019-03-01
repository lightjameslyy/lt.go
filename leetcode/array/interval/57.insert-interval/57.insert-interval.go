package main

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 *
 * https://leetcode.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (30.72%)
 * Total Accepted:    164.2K
 * Total Submissions: 534.6K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * Given a set of non-overlapping intervals, insert a new interval into the
 * intervals (merge if necessary).
 *
 * You may assume that the intervals were initially sorted according to their
 * start times.
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with
 * [3,5],[6,7],[8,10].
 *
 */
/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
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
思路：
	1. 找到newInterval.Start的index
	2. 把intervals[index]的end置为max(newInterval.End, intervals[index])
	3. 从下一个位置开始遍历，直到满足intervals[i].Start > intervals[i-1].End
*/
func insert(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	pos := -1
	// 找到插入的位置
	for i := 0; i < len(intervals); i++ {
		if newInterval.End < intervals[i].Start {
			// 在左侧，直接返回
			pos = i
			res = append(res, newInterval)
			for ; i < len(intervals); i++ {
				res = append(res, intervals[i])
			}
			return res
		} else if newInterval.Start <= intervals[i].End &&
			newInterval.End >= intervals[i].Start {
			// 相交，找到插入位置
			pos = i
			intervals[i].Start = min(intervals[i].Start, newInterval.Start)
			intervals[i].End = max(intervals[i].End, newInterval.End)
			break
		} else {
			// 在右边，加入当前的Interval，继续找pos
			res = append(res, intervals[i])
		}
	}
	// 没找到pos，说明已经找到最右边，追加newInterval，直接返回
	if pos == -1 {
		res = append(res, newInterval)
		return res
	}
	// merge interval
	res = append(res, intervals[pos])
	back := len(res) - 1
	for i := pos + 1; i < len(intervals); i++ {
		if intervals[i].Start > res[back].End {
			for ; i < len(intervals); i++ {
				res = append(res, intervals[i])
			}
			break
		} else {
			res[back].End = max(intervals[i].End, res[back].End)
		}
	}
	return res
}
