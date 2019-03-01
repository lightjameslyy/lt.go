package main

/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 *
 * https://leetcode.com/problems/game-of-life/description/
 *
 * algorithms
 * Medium (43.29%)
 * Total Accepted:    100.4K
 * Total Submissions: 231.8K
 * Testcase Example:  '[[0,1,0],[0,0,1],[1,1,1],[0,0,0]]'
 *
 * According to the Wikipedia's article: "The Game of Life, also known simply
 * as Life, is a cellular automaton devised by the British mathematician John
 * Horton Conway in 1970."
 *
 * Given a board with m by n cells, each cell has an initial state live (1) or
 * dead (0). Each cell interacts with its eight neighbors (horizontal,
 * vertical, diagonal) using the following four rules (taken from the above
 * Wikipedia article):
 *
 *
 * Any live cell with fewer than two live neighbors dies, as if caused by
 * under-population.
 * Any live cell with two or three live neighbors lives on to the next
 * generation.
 * Any live cell with more than three live neighbors dies, as if by
 * over-population..
 * Any dead cell with exactly three live neighbors becomes a live cell, as if
 * by reproduction.
 *
 *
 * Write a function to compute the next state (after one update) of the board
 * given its current state. The next state is created by applying the above
 * rules simultaneously to every cell in the current state, where births and
 * deaths occur simultaneously.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [0,1,0],
 * [0,0,1],
 * [1,1,1],
 * [0,0,0]
 * ]
 * Output:
 * [
 * [0,0,0],
 * [1,0,1],
 * [0,1,1],
 * [0,1,0]
 * ]
 *
 *
 * Follow up:
 *
 *
 * Could you solve it in-place? Remember that the board needs to be updated at
 * the same time: You cannot update some cells first and then use their updated
 * values to update other cells.
 * In this question, we represent the board using a 2D array. In principle, the
 * board is infinite, which would cause problems when the active area
 * encroaches the border of the array. How would you address these problems?
 *
 *
 */

/*
思路：
由于题目中要求我们用置换方法in-place来解题，所以我们就不能新建一个相同大小的数组，
那么我们只能更新原有数组，但是题目中要求所有的位置必须被同时更新，但是在循环程序
中我们还是一个位置一个位置更新的，那么当一个位置更新了，这个位置成为其他位置的
neighbor时，我们怎么知道其未更新的状态呢，我们可以使用状态机转换：
- 状态0： 死细胞转为死细胞
- 状态1： 活细胞转为活细胞
- 状态2： 活细胞转为死细胞
- 状态3： 死细胞转为活细胞
最后我们对所有状态对2取余，那么状态0和2就变成死细胞，状态1和3就是活细胞，达成目
的。
我们先对原数组进行逐个扫描，对于每一个位置，扫描其周围八个位置，如果遇到状
态1或2，就计数器累加1，扫完8个邻居，如果少于两个活细胞或者大于三个活细胞，
而且当前位置是活细胞的话，标记状态2，如果正好有三个活细胞且当前是死细胞的话，
标记状态3。完成一遍扫描后再对数据扫描一遍，对2取余变成我们想要的结果。
*/

func countAround(board [][]int, m, n, i, j int) (count int) {
	for ii := i - 1; ii <= i+1; ii++ {
		for jj := j - 1; jj <= j+1; jj++ {
			if ii == i && jj == j {
				continue
			}
			if ii >= 0 && ii < m && jj >= 0 && jj < n {
				if board[ii][jj] == 1 || board[ii][jj] == 2 {
					count++
				}
			}
		}
	}
	return
}

func gameOfLife(board [][]int) {
	const (
		D2D = iota // death to death
		L2L        // live to live
		L2D        // live to death
		D2L        // death to live
	)
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := countAround(board, m, n, i, j)
			if board[i][j] == 1 {
				if count < 2 || count > 3 {
					board[i][j] = 2
				}
			} else {
				if count == 3 {
					board[i][j] = 3
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] %= 2
		}
	}
	return
}
