<!-- .slide: data-background="img/PRIME_CHECKER/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Prime checker
In this task, you will learn about writing tests in Go by developing a prime tester in TDD-style.

----

### Too fast? Find source code here:
* [github.com/go-gurus/go_tour_src/tree/main/prime-checker](https://github.com/go-gurus/go_tour_src/tree/main/prime-checker)

----
<!-- .slide: data-background="img/PRIME_CHECKER/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

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
import "testing"

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

### Fake Test Data

* now lets fake some test data for our tests

----

### Too fast? Find source code here:
* [github.com/go-gurus/go_tour_src/tree/main/test-faker](https://github.com/go-gurus/go_tour_src/tree/main/test-faker)

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

* generate a new folder
* init a new module

```bash
go mod init grohm.io/test-data-faker
```

* add the package to the module

```bash
go get github.com/brianvoe/gofakeit/v6
```

----

* add a new `main.go` file

```go
// main.go
package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Email())
	fmt.Println(gofakeit.Phone())
	fmt.Println(gofakeit.BS())
	fmt.Println(gofakeit.BeerName())
	fmt.Println(gofakeit.Color())
	fmt.Println(gofakeit.Company())
	fmt.Println(gofakeit.HackerPhrase())
	fmt.Println(gofakeit.JobTitle())
	fmt.Println(gofakeit.CurrencyShort())
}
```
----

* lets run the code

```bash
go run main.go
```

```bash
Jessie Schowalter
billieruecker@dickens.info
8837610879
world-class
Double Bastard Ale
DarkGoldenRod
SlashDB
I'll compile the online USB transmitter, that should format the IB feed!
Executive
KRW
```
----

* now lets fake some credit card data

```go
// main.go
// ...
func main(){
	//...

	fmt.Println(gofakeit.CreditCardType())
	fmt.Println(gofakeit.CreditCardNumber(nil))
}
```

----

* refactor creditcard example

```go
// main.go
// ...
func main(){
	//...

	gofakeit.Seed(0)
	ccInfo := gofakeit.CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)
}
```
----

* lets run the code again

```bash
go run main.go
```

```bash
Shea Schinner
briannelegros@kutch.com
7854232973
rich
90 Minute IPA
SkyBlue
Farmers
Use the solid state HDD card, then you can transmit the open-source matrix!
Strategist
CVE
Hipercard
379037107654647
JCB
6443852913878647
06/28
445
```

----

* see the full list of fake data 
  * [github.com/brianvoe/gofakeit#functions](https://github.com/brianvoe/gofakeit#functions)

----

### What we have learned
* How to use the Go testing package
* How to run tests
* Write our first Go function with in and out parameters
* Basic loops and branching
* How to show code coverage
* How to create coverage reports
* How to fake test data
----

### Further readings
* testing
  * [pkg.go.dev/testing](https://pkg.go.dev/testing)
  * [go.dev/doc/tutorial/add-a-test](https://go.dev/doc/tutorial/add-a-test)
* loops
  * [go.dev/tour/flowcontrol/1](https://go.dev/tour/flowcontrol/1)
  * [geeksforgeeks.org/loops-in-go-language/](https://www.geeksforgeeks.org/loops-in-go-language/)
* gofakeit
  * [github.com/brianvoe/gofakeit](https://github.com/brianvoe)
  * [github.com/brianvoe/gofakeit#functions](https://github.com/brianvoe/gofakeit#functions)

---