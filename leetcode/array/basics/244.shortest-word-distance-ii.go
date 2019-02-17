// This is a follow up of Shortest Word Distance. The only difference is now you are given the list of words and your method will be calledrepeatedly many times with different parameters. How would you optimize it?
//
// Design a class which receives a list of words in the constructor, and implements a method that takes two words word1 and word2 and return the shortest distance between these two words in the list.
//
// For example,
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Given word1 = “coding”, word2 = “practice”, return 3.
// Given word1 = "makes", word2 = "coding", return 1.
//
// Note:
// You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.

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

type WordDistance struct {
	wdMap map[string][]int
}

func NewWordDistance(words []string) *WordDistance {
	res := new(WordDistance)
	wdMap := make(map[string][]int)
	for i, word := range words {
		if list, ok := wdMap[word]; ok {
			wdMap[word] = append(list, i)
		} else {
			wdMap[word] = []int{i}
		}
	}
	res.wdMap = wdMap
	return res
}

// 0 4 6
//  2   7
func (w *WordDistance) Shortest(word1, word2 string) int {
	res := math.MaxInt32
	list1, ok := w.wdMap[word1]
	if !ok {
		return res
	}
	list2, ok := w.wdMap[word2]
	if !ok {
		return res
	}
	i, j := 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			res = min(list2[j]-list1[i], res)
			i++
		} else {
			res = min(list1[i]-list2[j], res)
			j++
		}
	}
	if i == len(list1) {
		for ; j < len(list2); j++ {
			res = min(res, max(list1[i-1], list2[j])-min(list1[i-1], list2[j]))
		}
	}
	if j == len(list2) {
		for ; i < len(list1); i++ {
			res = min(res, max(list1[i], list2[j-1])-min(list1[i], list2[j-1]))
		}
	}
	return res
}

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1, word2 := "coding", "makes"
	// word1, word2 := "coding", "practice"
	wordDistance := NewWordDistance(words)
	fmt.Println(wordDistance.Shortest(word1, word2))
}
