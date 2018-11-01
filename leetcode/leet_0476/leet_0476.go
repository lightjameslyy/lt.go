package leet_0476

func FindComplement(num int) int {
	var res int
	factor := 1
	for num != 0 {
		remainder := num % 2
		num /= 2
		if remainder == 0 {
			res |= factor
		}
		factor = factor << 1
	}
	return res
}
