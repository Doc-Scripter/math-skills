package main

import (
	"fmt"
	"math"
	"os"

	o "math-skills/operations"
)

func main() {
	input := "data.txt"

	// check the arguments format
	if len(os.Args) <= 1 || os.Args[1] == "" || len(os.Args[2:]) != 0 {
		fmt.Println("Usage := go run . data.txt")
		os.Exit(0)y
	}
	if os.Args[1] != "data.txt" {
		fmt.Println("Application only accepts file data.txt")
		os.Exit(0)
	}

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("can not read file:", err)
	}

	lines := string(file)

	// check the input data
	o.CheckInput(lines)

	// mathematical operations
	Average := o.Average(lines)
	Median := o.Median(lines)
	Variance := o.Variance(lines)
	StandardDev := o.StandardDev(lines)

	// check if no overflow in output
	Max := float64(math.MaxInt64)
	Min := float64(math.MinInt64)
	if Average > Max || Median > Max || Variance > Max || StandardDev > Max {
		fmt.Println("the output exceeds capacity")
		os.Exit(0)
	}

	if Average <= Min || Median <= Min || Variance <= Min || StandardDev <= Min {
		fmt.Println("the output exceeds capacity")
		os.Exit(0)
	}

	// round to the nearest interger
	fmt.Println("Average:", int(math.Round(Average)))
	fmt.Println("Median:", int(math.Round(Median)))
	fmt.Println("Variance: ", int(math.Round(Variance)))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDev)))
}
