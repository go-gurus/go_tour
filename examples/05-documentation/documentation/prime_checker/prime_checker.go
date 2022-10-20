// Package prime_checker provides a prime number check function.
package prime_checker

import (
	"math"
)

// IsPrime check if int value is a prime number.
// It returns a boolean, true if it is prime number, false if not.
func IsPrime(value int) (result bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
