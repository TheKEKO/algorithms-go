// https://www.codewars.com/kata/57eaeb9578748ff92a000009/train/go
// Sum Mixed Array
package sum_mixed_array

import "strconv"

func SumMix(arr []any) int {
	sum := 0
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			sum += val
		case string:
			num, _ := strconv.Atoi(val)
			sum += num
		}
	}
	return sum
}
