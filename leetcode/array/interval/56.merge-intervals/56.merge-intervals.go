package main

// import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (34.79%)
 * Total Accepted:    306.4K
 * Total Submissions: 880.5K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 *
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 */
/**
 * Definition for an interval.
  type Interval struct {
	Start int
	End   int
}
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	pos := 0
	// 先按照Start排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= intervals[pos].End { // 相交
			intervals[pos].End = max(intervals[i].End, intervals[pos].End)
		} else {
			pos++
			if pos < i {
				intervals[pos] = intervals[i]
			}
		}
	}
	return intervals[:pos+1]
}
