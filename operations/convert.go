package operations

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Convert(s string) []float64 {
	var numbr []float64

	numStr := strings.Fields(s)
	for i := range numStr {
		num, err := strconv.ParseFloat(numStr[i], 64)
		if err != nil {
			fmt.Println("Can not convert", err)
			os.Exit(0)
		}
		numbr = append(numbr, num)

	}
	return numbr
}
