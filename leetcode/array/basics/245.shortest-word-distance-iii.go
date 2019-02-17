// This is a follow up of Shortest Word Distance. The only difference is now word1 could be the same asword2.
//
// Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.
//
// word1 and word2 may be the same and they represent two individual words in the list.
//
// For example,
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Given word1 = “makes”, word2 = “coding”, return 1.
// Given word1 = "makes", word2 = "makes", return 3.
//
// Note:
// You may assume word1 and word2 are both in the list.

package main

import (
	"fmt"
	"math"
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

func shortestDistance(words []string, word1, word2 string) int {
	prev1, prev2 := -1, -1 // 上一个word1和word2出现的位置
	res := math.MaxInt32
	// 如果word1==word2，只比较每两个word之间的距离即可
	if word1 == word2 {
		for i := 0; i < len(words); i++ {
			if words[i] == word1 {
				if prev1 != -1 {
					res = min(res, i-prev1)
				}
				prev1 = i
			}
		}
		return res
	}
	// 同243
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			if prev2 != -1 {
				res = min(res, i-prev2)
			}
			prev1 = i
		}
		if words[i] == word2 {
			if prev1 != -1 {
				res = min(res, i-prev1)
			}
			prev2 = i
		}
	}
	return res
}

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1, word2 := "makes", "makes"
	// word1, word2 := "coding", "practice"
	fmt.Println(shortestDistance(words, word1, word2))
}
