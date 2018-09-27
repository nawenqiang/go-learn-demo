package simplemath

import "math"

func Sqrt(val int) int {
	return int(math.Sqrt(float64(val)))
}
