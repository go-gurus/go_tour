package main

import (
	"math"
)

func main() {
}

func IsPrime(value int) (result bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
