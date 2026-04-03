// https://www.codewars.com/kata/55dc4520094bbaf50e0000cb/train/go
// Wilson primes

package wilson_primes

func AmIWilson(n int) bool {
	if n == 5 || n == 13 || n == 563 {
		return true
	}
	return false
}
