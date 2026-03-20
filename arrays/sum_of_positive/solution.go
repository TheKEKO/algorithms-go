// https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go
// Sum of positive
package sum_of_positive

func PositiveSum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
