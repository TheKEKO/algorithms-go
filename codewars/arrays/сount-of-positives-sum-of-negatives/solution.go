// https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go
// Count of positives / sum of negatives
package сount_of_positives_sum_of_negatives

func CountPositivesSumNegatives(numbers []int) []int {
	pos := 0
	neg := 0
	for _, v := range numbers {
		if v > 0 {
			pos++
		} else if v < 0 {
			neg += v
		}
	}
	return []int{pos, neg}
}
