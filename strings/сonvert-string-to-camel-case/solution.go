// https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go
package сonvert_string_to_camel_case

import (
	"unicode"
)

func ToCamelCase(s string) string {
	makeUpper := false
	result := ""
	for _, r := range s {
		if r == '-' || r == '_' {
			makeUpper = true
			continue
		}
		if makeUpper {
			r = unicode.ToUpper(r)
			makeUpper = false
		}
		result += string(r)
	}

	return result
}
