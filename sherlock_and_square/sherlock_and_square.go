package sherlockandsquare

import "math"

func SherlockAndSquare(a, b int32) int32 {

	lowerBound := math.Ceil(math.Sqrt(float64(a)))
	upperBound := math.Floor(math.Sqrt(float64(b)))

	delta := int32(upperBound - lowerBound)
	return delta + 1
}
