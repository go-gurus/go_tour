## generics

We want to create a go program that provides various functions for processing lists of different types.
Then we will optimize this program with generics and refactor the functions to show their advantages.

----

* first create a function that returns the sum of an integer map

```go
// main.go
package main

import "fmt"

func SumInts(m map[string]int64) int64 {
  var sum int64
  for _, val := range m {
    sum += val
  }
  return sum
}

```

----

* now lets create a similar function that returns the sum of an float map

```go
// main.go
// ...
func SumFloats(m map[string]float64) float64 {
  var sum float64
  for _, val := range m {
    sum += val
  }
  return sum
}
```

----

* now lets create the main function, create two maps and call the two functions

```go
// main.go
// ...
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
```

----

* execute the main function

```shell
$ go run main.go
Non-Generic Sums: 46 and 62.97
```

----

### Refactoring
Now lets switch to a generic function.

----

* add a function that can handle multiple types

```go
package main

import "fmt"

func SumIntsOrFloats[CompType comparable, ValueType int64 | float64](m map[CompType]ValueType) ValueType {
  var sum ValueType
  for _, val := range m {
    sum += val
  }
  return sum
}
```

----

* what is new here?
```
[CompType comparable, ValueType int64 | float64]
```
* the possible types for the function are defined
* the rest of the function uses the generic types
```go
func ... (m map[CompType]ValueType) ValueType {
```
```go
var sum ValueType
```

----

* lets create a `main` function again

```go
// main.go
// ...
func main() {
  ints := map[string]int64{
    "1st":  34,
	"2nd": 12,
  }

  floats := map[string]float64{
    "1st":  35.98,
	"2nd": 26.99,
  }

  fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))
}
```

----

* run the code

```shell
$ go run main.go
Generic Sums: 46 and 62.97
```

----

### Continue Refactoring
Now lets add some minor improvements to the code.

----

* lets add a type constraint

```go
// main.go
// ...
type Number interface {
  int64 | float64
}
```

----

* use the type constraint

```go
// main.go
// ...
func SumNumbers[CompType comparable, Number int64 | float64](m map[CompType]Number) Number {
  var sum Number
  for _, val := range m {
    sum += val
  }
  return sum
}
```

----

* adapt the main function, notice the `CompType` is not needed in the function call

```go
// main.go
// ...
func main() {

  fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
}
```

```shell
$ go run main.go
Generic Sums with Constraint: 46 and 62.97
```

----

### generic library lo
[github.com/samber/lo](https://github.com/samber/lo) contains many helpers for:
* slices
  * Filter, Map, ForEach, Subset, ...
* maps
  * Keys, Invert, ...
* strings
  * Substring, ...
* channels
  * SliceToChannel, Batch, ChannelMerge, ...
* many more
  

----

### What we have learned
* How to use maps
* How to use generics
* have a look on generic library

----
### Further readings
* generics tutorial
  * [go.dev/doc/tutorial/generics](https://go.dev/doc/tutorial/generics)
* lo - generics library
  * [pkg.go.dev/github.com/samber/lo](https://pkg.go.dev/github.com/samber/lo)
  * [github.com/samber/lo](https://github.com/samber/lo)

----
