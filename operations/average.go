package operations

import (
	"fmt"
	"strconv"
	"strings"
)

func Average(s string) float64 {
	var result float64

	numStr := strings.Fields(s)
	for _, v := range numStr {
		numbr, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("can not convert:", err)
		}

		result += (numbr)
	}

	finalresult := result / float64(len(numStr))
	return finalresult
}
