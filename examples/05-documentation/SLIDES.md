<!-- .slide: data-background="img/DOCUMENTATION/00.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

## Documentation
In this task, we want to build and serve some golang documentation.

----

<!-- .slide: data-background="img/DOCUMENTATION/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

### Create a module

* lets create a new folder and init a module

```shell
$ mkdir documentation
$ cd documentation
$ go mod init codecentric.de/documentation
```

----

* lets create a file `main.go` with some documentation

```go
// Package documentation provides a prime number check function and is a documentation showcase.
package documentation

func main() {
}
```

----

* lets build a new folder and file `prime_checker/prime_checker.go`

```go
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

```

----

* install the godoc tool

```shell
$ go install -v golang.org/x/tools/cmd/godoc@latest
```

* serve the documentation

```shell
$ ~/go/bin/godoc -http :6060
```

* visit [localhost:6060/pkg/codecentric.de/documentation/](http://localhost:6060/pkg/codecentric.de/documentation/)

----

* now lets add an other file `prime_checker/prime_checker_example_test.go`

```go
package prime_checker

import "fmt"

func ExampleIsPrime() {
	res := IsPrime(7)
	fmt.Println(res)
	//Output: true
}
```

----

* execute test case, change `//Output: false` and try again

```shell
$ go test ./...
```

* serve the documentation, check the example

```shell
$ ~/go/bin/godoc -http :6060
```

* visit [localhost:6060/pkg/codecentric.de/documentation/](http://localhost:6060/pkg/codecentric.de/documentation/)

----

### What we have learned
* How to write documentation for golang
* How to serve documentation for golang
* How to add example testcases into the documentation
----

### Further readings
* [pkg.go.dev/golang.org/x/tools/cmd/godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

---
