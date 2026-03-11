// https://www.codewars.com/kata/596fba44963025c878000039/train/go
package contamination__1_string

func Contamination(text, char string) string {
	if text == "" || char == "" {
		return ""
	}
	result := ""
	for i := 1; i <= len(text); i++ {
		result += char
	}
	return result
}
