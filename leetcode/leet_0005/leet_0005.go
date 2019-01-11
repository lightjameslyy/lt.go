package leet_0005

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

// 1. 贪心：如果当前的最长回文子串长度为maxlen，那么下一步只考虑长度为maxlen+1的子串
// Test Passed
func LongestPalindrome1(s string) string {
	l, r := 0, 0
	maxlen := 0
	for i := 0; i < len(s); i++ {
		for j := i + maxlen + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				maxlen = j - i
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
func LongestPalindrome2(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	maxlen := 0
	// 注意i和j的顺序
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			dp[i][j] = (s[i] == s[j]) && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > maxlen {
				maxlen = j - i + 1
				res = s[i : j+1]
			}
		}
	}
	return res
}

/*
3. 中心扩散法

helper(i, j)的定义：从[i, j]区间向两边试探

从中间向两侧试探，分奇偶讨论
    奇：helper(i, i)
    偶：helper(i, i+1)
*/

var res = ""

func helper(s string, l, r int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	if r-l-1 > len(res) {
		res = s[l+1 : r]
	}
}

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
    res = ""
	for i := 0; i < len(s); i++ {
		helper(s, i, i)
		helper(s, i, i+1)
	}
	return res
}
