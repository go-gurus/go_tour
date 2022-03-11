## Prime checker
In this task, you will learn about writing tests in Go by developing a prime tester in TDD-style.

----

### Solution

Lets write a test to specify what we expect from the prime checker yet to be developed.

```golang
package main
import 	"testing"

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
```

----
### Executing tests

```bash
go test main.go main_test.go

./main_test.go:9:6: undefined: IsPrime
./main_test.go:15:6: undefined: IsPrime
./main_test.go:35:14: undefined: IsPrime
FAIL    command-line-arguments [build failed]
FAIL
```

The test will fail, because the function `IsPrime` is yet to be implemented.

> We will learn a better way to run Go tests later 

----

### Implement the (naive) IsPrime function

```golang
func IsPrime(value int) (result bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value %i == 0 {
			return false
		}
	}
	return value > 1
}
```

Now,the tests run:

```bash
go test main.go main_test.go
ok      command-line-arguments  0.102s
```
----
### Table-driven tests

Cut down redundancy in your tests:
```golang
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
				t.Fatalf("Unexpected Result input=%d expected=%t actual=%t", 
					e.input, 
					e.expectedResult, 
					result)
			}
		})
	}
}
```

----
### What we have learned
* How to use the Go testing package
* How to run tests
* Write our first Go function with in and out parameters
* Basic loops and branching

---