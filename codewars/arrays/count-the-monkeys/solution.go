// https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/train/go
// Count the Monkeys!
package count_the_monkeys

func monkeyCount(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		res = append(res, i)
	}
	return res
}
