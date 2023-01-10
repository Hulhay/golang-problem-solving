package timconversion

import (
	"fmt"
	"strconv"
)

func TimeConversion(time string) string {

	timeType := time[len(time)-2:]
	hourStr := time[0:2]
	timeExcludeHour := time[2 : len(time)-2]

	hour, _ := strconv.Atoi(hourStr)
	if timeType == "PM" {
		hour += 12
		if hour >= 24 {
			hour -= 12
		}
	}

	if timeType == "AM" && hour == 12 {
		hour = 0
	}

	hourStr = strconv.Itoa(hour)
	if len(hourStr) == 1 {
		hourStr = fmt.Sprintf("0%s", hourStr)
	}

	return fmt.Sprintf("%s%s", hourStr, timeExcludeHour)
}
