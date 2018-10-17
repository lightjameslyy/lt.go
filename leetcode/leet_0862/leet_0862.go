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
		for r < lenA {
			cursum += A[r]
			if cursum <= 0 {
				l = r + 1
			} else if cursum >= K {
				if r-l+1 < curlen {
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
