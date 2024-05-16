package operations

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// function to check if there is overflow in the input data
func Overflow(s string) (bool, error, []float64) {
	numStr := strings.Fields(s)
	var IsTrue bool
	for _, v := range numStr {
		var numbers []float64
		numbr, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("the data is not a number")
			os.Exit(0)
		}

		if numbr > math.MaxInt64 {
			IsTrue = true
		}
		if numbr < math.MinInt64 {
			IsTrue = true
		}
		numbers = append(numbers, numbr)
		return false, nil, numbers
	}
	return IsTrue, nil, []float64{}
}
