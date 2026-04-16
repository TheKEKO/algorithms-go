// https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/go
// Find the smallest integer in the array
package find_the_smallest_integer_in_the_array

import "sort"

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
