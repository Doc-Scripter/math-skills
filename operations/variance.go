package operations

import (
	"strings"
)

func Variance(s string) float64 {
	average := Average(s)

	// var num []float64
	var result []float64
	var finalresult float64

	numStr := strings.Fields(s)
	num := Convert(s)
	// calculate the variance
	for i := range num {
		v := average - num[i]
		result = append(result, v*v)
	}

	for i := range result {
		finalresult += result[i]
	}
	// minus 1 to avoid biasness in the calculation
	finalresult = finalresult / float64(len(numStr))

	return finalresult
}
