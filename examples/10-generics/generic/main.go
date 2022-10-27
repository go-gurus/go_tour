package main

import "fmt"

func SumIntsOrFloats[CompType comparable, ValueType int64 | float64](m map[CompType]ValueType) ValueType {
	var sum ValueType
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

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
}
