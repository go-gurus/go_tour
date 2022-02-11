package main

import (
	"testing"
)

func TestPrimeCheckerTheNaiveWay(t *testing.T) {
	t.Run("should return FALSE when no prime number given", func(t *testing.T) {
		if IsPrime(4) == true {
			t.Fatal("Reported IsPrime=true for 4")
		}
	})

	t.Run("should return TRUE when prime number given", func(t *testing.T) {
		if IsPrime(7) == false {
			t.Fatal("Reported IsPrime=true for 7")
		}
	})
}

func TestPrimeCheckerTableDriven(t *testing.T) {
	cases := []struct {
		input          int
		expectedResult bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{7, true},
	}

	for _, e := range cases {
		t.Run("Should return expected result ", func(t *testing.T) {
			result := IsPrime(e.input)
			if result != e.expectedResult {
				t.Fatalf("Unexpected Result input=%d expected=%t actual=%t", e.input, e.expectedResult, result)
			}
		})
	}
}
