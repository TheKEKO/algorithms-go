// https://www.codewars.com/kata/5a2fd38b55519ed98f0000ce/train/go
package multiplication_table_for_number

import "fmt"

func MultiTable(number int) string {
	result := ""
	for i := 1; i <= 10; i++ {
		sum := number * i
		if i == 10 {
			result += fmt.Sprintf("%d * %d = %d", i, number, sum)
		} else {
			result += fmt.Sprintf("%d * %d = %d\n", i, number, sum)
		}
	}
	return result
}
