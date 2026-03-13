// https://www.codewars.com/kata/56cd44e1aa4ac7879200010b/train/go
package is_the_string_uppercase

import "unicode"

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, v := range s {
		if unicode.IsLetter(v) && !unicode.IsUpper(v) {
			return false
		}
	}
	return true
}
