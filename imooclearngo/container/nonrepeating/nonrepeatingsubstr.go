package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr1(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 支持中文
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"),
		lengthOfNonRepeatingSubStr("bbbb"),
		lengthOfNonRepeatingSubStr("pwwkew"),
		lengthOfNonRepeatingSubStr(""),
		lengthOfNonRepeatingSubStr("b"),
		lengthOfNonRepeatingSubStr("abcdef"),
		lengthOfNonRepeatingSubStr("中文测试"),
		lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞"),
	)
}
