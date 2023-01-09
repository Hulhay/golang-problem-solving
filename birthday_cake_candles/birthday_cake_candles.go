package birthdaycakecandles

func BirthdayCakeCandles(candles []int32) int32 {

	tallest := candles[0]
	count := 1

	for i := 1; i < len(candles); i++ {
		if candles[i] == tallest {
			count++
		}
		if candles[i] > tallest {
			tallest = candles[i]
			count = 1
		}
	}

	return int32(count)
}
