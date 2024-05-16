package operations

import (
	"fmt"
	"os"
	"strings"
)

// function to check input file
func CheckInput(s string) []float64{
	numStr := strings.Fields(s)

	// check if there is overflow in input data
	v, err,numbers := Overflow(s)
	if err != nil {
		fmt.Println("the data  is not a number")
		os.Exit(0)
	}
	if v == true {
		fmt.Println("the data exceeds capacity")
		os.Exit(0)
	}
	line := strings.Split(s, "\n")

	// check if the data has only one value
	if len(numStr) == 1 {
		fmt.Println("Since this is a population statistical measure scenario please use more than one dataset e.g 21 22")
		os.Exit(0)
	}
	// check format of data
	for i := range line {
		numbrcount := strings.Fields(line[i])
		if len(numbrcount) > 1 {
			fmt.Println("the format for data in data.txt should be : \n number \n number \n ...")
			os.Exit(0)
		}
	}

	// check if the file is empty
	if s == "" {
		fmt.Println("data.txt is empty")
		os.Exit(0)
	}

	var Totals string
	for i := 0; i < len(line)-1; i++ {
		if len(line[i]) > 0 {
			Totals = line[i] + line[i+1]
		}
	}
	if len(Totals) == 0 {
		fmt.Println("data.txt is empty")
		os.Exit(0)

	}
	return numbers
}
