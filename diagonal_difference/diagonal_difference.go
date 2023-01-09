package diagonaldifference

import "math"

func DiagonalDifference(arr [][]int) int {

	size := len(arr)

	for i := 0; i < size; i++ {
		if len(arr[i]) != size {
			return 0
		}
	}

	sumPrimary := 0
	sumSecondary := 0

	for i := 0; i < size; i++ {
		sumPrimary += arr[i][i]
		sumSecondary += arr[i][size-(i+1)]
	}

	return int(math.Abs(float64(sumPrimary) - float64(sumSecondary)))
}
