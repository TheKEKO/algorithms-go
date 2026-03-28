// https://www.codewars.com/kata/583710ccaa6717322c000105/train/go
// Simple multiplication

package simple_multiplication

func SimpleMultiplication(n int) int {
	if n%2 == 0 {
		return n * 8
	} else {
		return n * 9
	}
}
