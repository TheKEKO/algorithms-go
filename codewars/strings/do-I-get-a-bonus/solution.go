// https://www.codewars.com/kata/56f6ad906b88de513f000d96
package strings

import (
	"strconv"
)

func BonusTime(salary int, bonus bool) string {
	if bonus {
		salary *= 10
	}
	return "£" + strconv.Itoa(salary)
}
