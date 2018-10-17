package leet_0862

func ShortestSubarray(A []int, K int) int {
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
		for r < lenA && l < lenA {
			curnum := A[r]
			if r-l == curlen {
				if curnum <= 0 {
					l = r + 1
					cursum = 0
				} else if curnum <= A[l] {
					cursum -= A[l]
					cursum -= A[l+1]
					l += 2
					for A[l] <= 0 {
						l++
					}
				} else {
					for curnum >= A[l] {
						curnum -= A[l]
						l++
					}
					if r-l+1 < curlen {
						curlen = r - l + 1
						if curlen == 2 {
							return 2
						}
					} else {
						cursum -= A[l]
						cursum -= A[l+1]
						l += 2
						for A[l] <= 0 {
							l++
						}
					}
				}
				continue
			}
			cursum += A[r]
			if cursum <= 0 {
				l = r + 1
				cursum = 0
			} else if cursum >= K {
				cursum = A[r]
				l = r
				for cursum < K {
					l--
					cursum += A[l]
				}
				if r-l+1 < curlen {
					curlen = r - l + 1
				}
				if curlen == 2 {
					return 2
				}
			} else {
				if r-l+1 >= curlen {
					cursum -= A[l]
					l++
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
