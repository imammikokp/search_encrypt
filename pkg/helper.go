package pkg

import (
	"math"
	"strconv"
)

func MedianRangeValue(min, max int) (leftMin, leftMax, rightMin, rightMax int) {
	MedianValue := math.RoundToEven(float64(max-min) / 2)
	leftMin = min
	leftMax = int(MedianValue) + min
	rightMin = int(MedianValue) + min + 1
	rightMax = max
	return
}

func StringToInt(value string) int {

	if value != "" {
		result, _ := strconv.Atoi(value)
		return result
	}
	return 0
}