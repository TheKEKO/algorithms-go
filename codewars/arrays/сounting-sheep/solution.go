// https://www.codewars.com/kata/54edbc7200b811e956000556/train/go
// Counting sheep...

package сounting_sheep

func CountSheeps(numbers []bool) int {
	res := 0
	for _, v := range numbers {
		if v {
			res++
		}
	}
	return res
}
