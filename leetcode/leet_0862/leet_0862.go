package leet_0862

// time limit exceeded
func ShortestSubarray1(A []int, K int) int {
	res := -1
	for _, val := range A {
		if val >= K {
			return 1
		}
	}
	if K > 0 {
		cursum := 0
		curlen := 50001 // A.length <= 50000
		var l, r int
		lenA := len(A)
		for r < lenA {
			cursum += A[r]
			if cursum <= 0 {
				l = r + 1
				cursum = 0
				r++
				continue
			}
			if cursum >= K {
				l = r
				tmpsum := 0
				for tmpsum < K {
					tmpsum += A[l]
					l--
				}
				l++
				if r-l+1 < curlen {
					curlen = r - l + 1
					if curlen == 2 {
						return 2
					}
				}
			}
			// find sub max sum as next candidate sequence
			if r-l+1 >= curlen {
				submax := 0
				subsum := 0
				p := r
				index := l
				for p > l {
					subsum += A[p]
					if subsum > submax {
						submax = subsum
						index = p
					}
					p--
				}
				cursum = submax
				l = index
				if cursum >= K && r-l+1 < curlen {
					curlen = r - l + 1
				}
			}
			r++
		}
		if curlen < 50001 {
			res = curlen
		}
	}
	return res
}

// wrong
func ShortestSubarray2(A []int, K int) int {
	lenA := len(A)
	min_len := lenA + 1
	cursum, start := 0, 0
	for i := 0; i < lenA; i++ {
		for cursum < K && i < lenA {
			cursum += A[i]
			i++
		}
		for cursum >= K && start < lenA {
			if i-start < min_len {
				min_len = i - start
			}
			cursum -= A[i]
			i++
		}
	}
	return min_len
}

// put it another way(dp):
// 	dp[0] = 0
// 	dp[i] = sum of A[0] to A[i-1]	i >= 1
// equivalent problem:
// 	find min(j-i) in dp such that:
// 		dp[j] - dp[i] >= K
func ShortestSubarray(A []int, K int) int {
	dp := make([]int, len(A)+1)
	for i, val := range A {
		if val >= K {
			return 1
		}
		dp[i+1] = dp[i] + val
	}

	res := len(dp)
	Q := []int{0}
	for i := 1; i < len(dp); i++ {
		for len(Q) > 0 && dp[i]-dp[Q[0]] >= K {
			if i-Q[0] < res {
				res = i - Q[0]
			}
			Q = Q[1:]
		}
		for len(Q) > 0 && dp[i] < dp[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, i)
	}

	if res == len(dp) {
		return -1
	}
	return res
}
