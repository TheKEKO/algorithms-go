// https://www.codewars.com/kata/5583090cbe83f4fd8c000051/train/go
// Convert number to reversed array of digits

package convert_number_to_reversed_array_of_digits

func Digitize(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var res []int
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	return res
}
