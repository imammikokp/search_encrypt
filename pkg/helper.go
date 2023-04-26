package pkg

import (
	"math"
)

func MedianRangeValue(min, max int) (leftMin, leftMax, rightMin, rightMax int) {
	MedianValue := math.RoundToEven(float64(max-min) / 2)
	leftMin = min
	leftMax = int(MedianValue) + min
	rightMin = int(MedianValue) + min + 1
	rightMax = max
	return
}
