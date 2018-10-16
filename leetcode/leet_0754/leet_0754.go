package leet_0754

/*
 * delta = (1 + 2 + ... + k) - abs(target)		// make k minimal
 * key point: delta even or odd
 * 		even: must exist subsum of (1, 2, ..., k) = k/2
 * 		odd: add k+1 or k+2 to make delta even
 */

func ReachNumber(target int) int {
	// only consider abs(target)
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	// now delta == abs(target)
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}
