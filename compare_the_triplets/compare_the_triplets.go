package comparethetriplets

func CompareTheTriplets(arrA, arrB []int) []int {

	result := []int{0, 0}

	if len(arrA) != len(arrB) {
		return result
	}

	for i := 0; i < len(arrA); i++ {
		if arrA[i] > arrB[i] {
			result[0]++
		}
		if arrA[i] < arrB[i] {
			result[1]++
		}
	}

	return result
}
