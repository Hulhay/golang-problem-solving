package minimaxsum

import (
	"math"
)

func MiniMaxSum(arr []int32) []int32 {

	min := math.Inf(1)
	max := math.Inf(-1)

	total := int32(0)
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	for i := 0; i < len(arr); i++ {
		if float64(total-arr[i]) > max {
			max = float64(total - arr[i])
		}
		if float64(total-arr[i]) < min {
			min = float64(total - arr[i])
		}
	}

	return []int32{int32(min), int32(max)}
}
