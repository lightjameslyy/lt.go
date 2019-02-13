// Shortest Word Distance
// Total Accepted: 1754 Total Submissions: 4239 Difficulty: Easy
// Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.
//
// For example,
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Given word1 = “coding”, word2 = “practice”, return 3.
// Given word1 = "makes", word2 = "coding", return 1.
//
// Note:
// You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.

package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// two pointers: O(n)
func shortestDistance(words []string, word1, word2 string) int {
	prev1, prev2 := -1, -1 // 上一个word1和word2出现的位置
	i := 0
	// 先找到第一次prev1和prev2都有效的位置
	for ; prev1 == -1 || prev2 == -1; i++ {
		if words[i] == word1 {
			prev1 = i
		}
		if words[i] == word2 {
			prev2 = i
		}
	}
	res := max(prev1, prev2) - min(prev1, prev2)
	// 更新res
	for ; i < len(words); i++ {
		if words[i] == word1 {
			res = min(res, i-prev2)
			prev1 = i
		}
		if words[i] == word2 {
			res = min(res, i-prev1)
			prev2 = i
		}
	}
	return res
}

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	//word1, word2 := "coding", "makes"
	word1, word2 := "coding", "practice"
	fmt.Println(shortestDistance(words, word1, word2))
}
