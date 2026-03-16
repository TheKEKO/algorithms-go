// https://www.codewars.com/kata/56453a12fcee9a6c4700009c/train/go
package compare_within_margin

import (
	"math"
)

func CloseCompare(a, b, margin float64) int {
	if math.Abs(a-b) <= margin {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
