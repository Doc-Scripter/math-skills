package operations

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// function to get the median
func Median(s string) float64 {
	var numbr []int

	numStr := strings.Fields(s)
	for i := range numStr {
		num, err := strconv.ParseFloat(numStr[i], 64)
		if err != nil {
			fmt.Println("Can not convert", err)
		}
		numbr = append(numbr, int(math.Round(num)))

	}
	// sort the numbers in order from least
	sort.Ints(numbr)

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
