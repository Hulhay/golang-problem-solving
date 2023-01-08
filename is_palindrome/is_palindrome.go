package ispalindrome

func IsPalindrome(word string) bool {
	reversed := ""

	i := len(word) - 1
	for i >= 0 {
		reversed += string(word[i])
		i--
	}

	return word == reversed
}

func IsPalindromeV2(word string) bool {

	isEqual := true
	for i := 0; i < len(word)/2; i++ {
		if string(word[i]) != string(word[len(word)-1-i]) {
			isEqual = false
			break
		}
	}

	return isEqual
}
