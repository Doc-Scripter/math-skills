package operations

import (
	"math"
	"strconv"
	"strings"
)

// function to check if there is overflow in the input data
func Overflow(s string) (bool, error) {
	numStr := strings.Fields(s)
	var IsTrue bool
	for _, v := range numStr {
		numbr, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return false, err
		}
		if numbr > math.MaxInt64 {
			IsTrue = true
		}
		if numbr < math.MinInt64 {
			IsTrue = true
		}
	}
	return IsTrue, nil
}
