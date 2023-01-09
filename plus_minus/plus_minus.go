package plusminus

import "fmt"

// return array of string that take float number with 6 decimal
func PlusMinus(arr []int32) []string {

	size := len(arr)
	positiveNumber := 0
	negativeNumber := 0
	zeroNumber := 0

	for _, num := range arr {
		if num > 0 {
			positiveNumber++
		}
		if num < 0 {
			negativeNumber++
		}
		if num == 0 {
			zeroNumber++
		}
	}

	result := []string{
		fmt.Sprintf("%.6f", float32(positiveNumber)/float32(size)),
		fmt.Sprintf("%.6f", float32(negativeNumber)/float32(size)),
		fmt.Sprintf("%.6f", float32(zeroNumber)/float32(size)),
	}

	return result
}
