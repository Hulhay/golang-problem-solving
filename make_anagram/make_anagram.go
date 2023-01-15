package makeanagram

import "math"

func MakeAnagram(string1, string2 string) int32 {

	mapping := map[rune]int{}

	for _, char := range string1 {
		mapping[char]++
	}

	for _, char := range string2 {
		mapping[char]--
	}

	var count int32
	for c := range mapping {
		if mapping[c] != 0 {
			count += int32(math.Abs(float64(mapping[c])))
		}
	}

	return count
}
