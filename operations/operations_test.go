package operations

import (
	"testing"
)

type operation struct {
	input    string
	expected float64
}

func TestAverage(t *testing.T) {
	op := operation{input: "1 2 3 4 5", expected: 3.0}
	mean := Average(op.input)

	if mean != op.expected {
		t.Errorf("Expected %f, but got %f", op.expected, mean)
	}
}

func TestMedian(t *testing.T) {
	op := operation{input: "1 2 3 4 5", expected: 3}
	median := Median(op.input)

	if median != op.expected {
		t.Errorf("Expected %f ,but got %f", op.expected, median)
	}
}

func TestVariance(t *testing.T) {
	op := operation{input: "1 2 3 4 5", expected: 2.5}
	variance := Variance(op.input)

	if variance != op.expected {
		t.Errorf("Expected %f but got %f", op.expected, variance)
	}
}

func TestStandardDev(t *testing.T) {
	op := operation{input: "1 2 3", expected: 1}
	standardDev := StandardDev(op.input)

	if standardDev != op.expected {
		t.Errorf("Expected %f but got :%f", op.expected, standardDev)
	}
}
