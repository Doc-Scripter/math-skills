package operations

import (
	"sort"
)

// function to get the median
func Median(s string) float64 {
	var numbr []float64

	for _, v := range Convert(s) {
		numbr = append(numbr, (v))
	}
	// sort the numbers in order from least
	sort.Float64s(numbr)

	// get the median when numbers are even or odd
	if len(numbr)%2 != 0 {
		return float64(numbr[(len(numbr) / 2)])
	} else {
		var float float64
		mid := len(numbr) / 2
		float = float64(numbr[mid]+numbr[mid-1]) / float64(2)
		return float
	}
}
