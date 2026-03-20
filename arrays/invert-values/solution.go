// https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/train/go
package invert_values

func Invert(arr []int) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = -v
	}
	return res
}
