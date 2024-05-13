// Convert a String to a Number!

// Solution 1
import "strconv"

func StringToNumber(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
