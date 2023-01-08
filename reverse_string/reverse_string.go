package reversestring

import "strings"

// resverse string and ignore whitespace, number, and punctuation
func ReverseString(word string) (reversed string) {
	reversed = ""
	i := len(word) - 1
	for i >= 0 {
		if int(word[i]) >= int('A') && int(word[i]) <= int('z') {
			reversed += strings.ToLower(string(word[i]))
		}
		i--
	}
	return strings.ToLower(reversed)
}
