package operations

func Average(s string) float64 {
	var result float64
	numbers := Convert(s)

	for i := range numbers {
		result += numbers[i]
	}

	finalresult := result / float64(len(numbers))
	return finalresult
}
