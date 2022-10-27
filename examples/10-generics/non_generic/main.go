package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, val := range m {
		sum += val
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, val := range m {
		sum += val
	}
	return sum
}

func main() {
	ints := map[string]int64{
		"1st": 34,
		"2nd": 12,
	}

	floats := map[string]float64{
		"1st": 35.98,
		"2nd": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))
}
