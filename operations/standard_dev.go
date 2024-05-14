package operations

import "math"

// standard deviation is the square root of variance
func StandardDev(s string) float64 {
	result := float64(Variance(s))
	finalresult := math.Sqrt(result)
	return finalresult
}
