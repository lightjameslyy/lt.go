package leet_0005

import "fmt"

/*
5. Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"

*/

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// 1. 贪心：如果当前的最长回文子串长度为maxLen，那么下一步只考虑长度为maxLen+1的子串
// Test Passed
func LongestPalindrome1(s string) string {
	l, r := 0, 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i + maxLen + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				maxLen = j - i
				l = i
				r = j
			}
		}
	}
	return s[l:r]
}

/*
2. DP
time: O(n^2) space: O(n^2)
两种基本形式：
	aba
	abba

*/
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	maxLen := 0
	// 注意i和j的顺序
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			dp[i][j] = (s[i] == s[j]) && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				res = s[i : j+1]
			}
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return res
}

/*
3. (TODO)中心扩散法
*/
