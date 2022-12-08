<!-- .slide: data-background="img/PRIME_CHECKER/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Prime checker
In this task, you will learn about writing tests in Go by developing a prime tester in TDD-style.

----

<!-- .slide: data-background="img/PRIME_CHECKER/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

### Solution

* lets write a simple `main.go` file

```go
// main.go
package main

func main() {
}
```

----

* now, Lets write a test to specify what we expect from the prime checker yet to be developed.

```golang
// main_test.go
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
			t.Fatal("Reported IsPrime=false for 7")
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
// main.go
// ...
import (
  "math"
)

func IsPrime(value int) (result bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value %i == 0 {
			return false
		}
	}
	return value > 1
}
// ...
```

* run the tests again

```bash
go test main.go main_test.go
ok      command-line-arguments  0.102s
```
----
### Table-driven tests

Cut down redundancy in your tests:
```golang
// main_test.go
// ...
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
### Code Coverage

* now lets execute the tests with coverage

```bash
$ go test -cover main.go main_test.go
ok      command-line-arguments  0.191s  coverage: 100.0% of statements
```
----

* lets add also a report

```bash
go test -coverprofile=coverage.out main.go main_test.go
```

```text
mode: set
/Users/grohmio/repos/cc/gophers/golang-for-developers/examples/03-prime-checker/main.go:7.14,8.2 0 0
/Users/grohmio/repos/cc/gophers/golang-for-developers/examples/03-prime-checker/main.go:10.39,11.67 1 1
/Users/grohmio/repos/cc/gophers/golang-for-developers/examples/03-prime-checker/main.go:16.2,16.18 1 1
/Users/grohmio/repos/cc/gophers/golang-for-developers/examples/03-prime-checker/main.go:11.67,12.19 1 1
/Users/grohmio/repos/cc/gophers/golang-for-developers/examples/03-prime-checker/main.go:12.19,14.4 1 1
```
----
* now lets use the go `coverage` tool to generate a graphical report

```bash
go tool cover -html=coverage.out main.go main_test.go
```

![go-playground](img/PRIME_CHECKER/02.png)<!-- .element height="500px" -->

----
### What we have learned
* How to use the Go testing package
* How to run tests
* Write our first Go function with in and out parameters
* Basic loops and branching
* How to show code coverage
* How to create coverage reports

---