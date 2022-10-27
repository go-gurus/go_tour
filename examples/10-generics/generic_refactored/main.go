package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumNumbers[CompType comparable, Number int64 | float64](m map[CompType]Number) Number {
	var sum Number
	for _, val := range m {
		sum += val
	}
	return sum
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
