// https://www.codewars.com/kata/51c8991dee245d7ddf00000e/train/go
// Reversed Words
package reversed_words

import "strings"

func ReverseWords(str string) string {
	words := strings.Fields(str)

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	return strings.Join(words, " ")
}
