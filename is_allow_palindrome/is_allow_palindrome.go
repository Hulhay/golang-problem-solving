package isallowpalindrome

func IsAllowPalindrome(s string) bool {

	mapping := map[string]int{}

	for _, val := range s {
		if _, ok := mapping[string(val)]; ok {
			mapping[string(val)] += 1
		} else {
			mapping[string(val)] = 1
		}
	}

	count := 0
	for _, n := range mapping {
		if n%2 != 0 {
			count += 1
		}
		if count > 1 {
			return false
		}
	}

	return true
}
