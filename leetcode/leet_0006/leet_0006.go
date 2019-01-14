package leet_0006

import "fmt"

/*
6. ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

*/

/*
input: P  A  Y  P  A  L  I  S  H  I  R  I  N  G
       0  1  2  3  4  5  6  7  8  9  10 11 12 13
先看一下两个例子：
1. numRows = 3
                        indexs in input             span
    P   A   H   N       0, 4, 8, 12                 4
    A P L S I I G       1, 3, 5, 7, 9, 11, 13       2
    Y   I   R           2, 6, 10                    4
2. numRows = 4
                        indexs in input             span
    P     I    N        0, 6, 12                    6
    A   L S  I G        1, 5, 7, 11, 13             4, 2, 4, 2
    Y A   H R           2, 4, 8, 10                 2, 4, 2
    P     I             3, 9                        6
3. numRows = 5
                        indexs in input             span
    P       H           0, 8                        8
    A     S I           1, 7, 9                     6, 2
    Y   I   R           2, 6, 10                    4, 4
    P L     I G         3, 5, 11, 13                2, 6, 2
    A       N           4, 12                       8

总结span的规律：
    1. span的最大值：span = numRows * 2 - 2
    2. 前numRows-1行：起始span递减2
    3. 第numRows行：span

*/

func Convert(s string, numRows int) string {
	if numRows < 2 || numRows >= len(s) {
		return s
	}
	res := make([]byte, len(s))
	resi := 0
	span, step := numRows*2-2, 0
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			step = span
		} else {
			step = i * 2
		}
		for j := i; j < len(s); {
			res[resi] = s[j]
			resi++
			if step != span {
				step = span - step
			}
			j = j + step
		}
	}
	fmt.Println(string(res))
	return string(res)
}
