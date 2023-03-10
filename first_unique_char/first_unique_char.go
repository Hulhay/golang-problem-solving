package firstuniquechar

func FirstUniqueChar(s string) int {

	m := [26]int{}

	for _, char := range s {
		m[char-'a'] += 1
	}
	for i, char := range s {
		if m[char-'a'] == 1 {
			return i
		}
	}
	return -1
}
