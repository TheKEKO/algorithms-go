// https://www.codewars.com/kata/5899642f6e1b25935d000161/train/go
// Merge two sorted arrays into one
package merge_two_sorted_arrays_into_one

import "sort"

func MergeArrays(arr1, arr2 []int) []int {
	merged := append(arr1, arr2...)
	sort.Ints(merged)

	if len(merged) == 0 {
		return merged
	}

	var res []int
	unique := make(map[int]bool)

	for _, v := range merged {
		if !unique[v] {
			unique[v] = true
			res = append(res, v)
		}
	}
	return res
}
