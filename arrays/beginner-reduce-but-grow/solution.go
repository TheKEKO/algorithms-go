// https://www.codewars.com/kata/57f780909f7e8e3183000078/train/go
package beginner_reduce_but_grow

func Grow(arr []int) int {
	res := 1
	for i := 0; i < len(arr); i++ {
		res *= arr[i]
	}
	return res
}
