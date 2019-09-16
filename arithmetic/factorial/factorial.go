package factorial

/**
 * @param n: A long integer
 * @return: An integer, denote the number of trailing zeros in n!
 */
func TrailingZeros(n int64) int64 {
	// write your code here, try to do it without arithmetic operators.
	var p, cnt int64
	p = n / 5
	for p > 0 {
		cnt += p
		p /= 5
	}
	return cnt
}
