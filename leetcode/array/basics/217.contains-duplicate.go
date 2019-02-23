package main

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (50.60%)
 * Total Accepted:    297.7K
 * Total Submissions: 588.3K
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 *
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: true
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 */

//import "fmt"

// input:
//      l: left index
//      r: right index
//      p: pivot index
func Partition(nums []int, l, r, p int) int {
	pivotVal := nums[p]
	// 把pivot移到最右边，然后处理nums[l, r-1]即可
	nums[p], nums[r] = nums[r], nums[p]
	np := l
	for i := l; i <= r-1; i++ {
		if nums[i] < pivotVal {
			nums[np], nums[i] = nums[i], nums[np]
			np++
		}
	}
	nums[np], nums[r] = nums[r], nums[np]
	// 返回pivot value最终的位置
	// fmt.Println(l, r, np)
	return np
}

func QuickSort(nums []int, l, r int) {
	if r > l {
		p := l + (r-l)/2
		np := Partition(nums, l, r, p)
		// fmt.Println(nums)
		QuickSort(nums, l, np-1)
		QuickSort(nums, np+1, r)
	}
}

// method 1: 先排序，然后遍历一边，检查是否有相邻元素相同
// 顺便练习一下快排
func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	QuickSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// method 2: hash，时间复杂度低，但是要O(n)的空间，略

//func main() {
//	nums := []int{2, 6, 5, 2, 3, 1, 7}
//	fmt.Println(nums)
//	QuickSort(nums, 0, len(nums)-1)
//	fmt.Println(nums)
//}
