// https://www.codewars.com/kata/5a00e05cc374cb34d100000d
package arrays

func ReverseSeq(n int) []int {
	arr := make([]int, n)

	for i := range arr {
		arr[i] = n - i
	}

	return arr
}
