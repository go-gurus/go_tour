<!-- .slide: data-background="img/INTRODUCTION/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

# Go Tour

----
## Authors

|                                                                |                                                                                                                                                                                                                                                                                    |
|----------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ![Manfred](img/cc_manfred.jpg)<!-- .element height="200px" --> | <ul><li>[Manfred Dreese](https://codecentric.de)<br />[dreese.de](https://dreese.de)<br />[github/i78](https://github.com/i78)<br />[twitter/dem9d](https://twitter.com/dem9d)<br /></li></ul>                                                                                     |
| ![grohmio](img/cc_grohmio.png)<!-- .element height="200px" --> | <ul><li>[Andreas Grohmann](https://grohm.io)<br />[grohm.io](https://grohm.io)<br />[github/grohmio](https://github.com/grohmio)<br />[twitter/grohmeo](https://twitter.com/grohmeo)<br />[stackoverflow/grohmio](https://stackoverflow.com/users/6654539/grohmio)</li></ul> |

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00-big.jpg" data-background-size="100%" data-background-position="50% 50%" -->

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

## Materials
* slides
  * [go-gurus.github.io/go_tour]( https://go-gurus.github.io/go_tour)
  * [go.grohm.io](https://go.grohm.io)
* source code
  * [github.com/go-gurus/go_tour_src](https://github.com/go-gurus/go_tour_src)
* presentation
  * [github.com/go-gurus/go_tour](https://github.com/go-gurus/go_tour)
* Slack
  * [go-gurus.slack.com](https://go-gurus.slack.com)
* github
  * [github.com/go-gurus](https://github.com/go-gurus)

----

## Workshop Abstract
Golang powers some of the most relevant IT projects of the last decade, such as Docker or Kubernetes. A lot of this success can be traced back to its simplicity, efficiency, tooling and developer experience.

In this workshop, we assume that the participants are already fluent with another programming language, such as Java, Python or C#. We will cover the go basics, walk along some real-world examples and build some REST/GraphQL and Data Analysis usecases.

---
<!-- .slide: data-background="img/INTRODUCTION/01.jpg" data-background-size="100%" data-background-position="50% 50%" -->

----

## Introduction
Why should you look at the GO programming language?
Go is an open source programming language. 
But go takes a lot of getting used to and is reduced to the essentials.
You can build anything from small helper scripts for pipelines to microservices to highly complex monoliths with GO.
The advantages of Go are a strong standardization of the tools,
a high applicability of the binaries and a high performance of the build process.
We will discuss further advantages of Go in more detail.

----

### Why Go?
* small binaries
* fast build times
* multi platform builds
* binaries could be used in scratch image
* small set of keywords (25)
* parallelism
* completely downward compatible

----
<!-- .slide: data-background="img/INTRODUCTION/go_projects.jpg" data-background-size="80%" data-background-position="50% 50%" -->

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

### Install Golang Linux, Windows, MacOS
* download your package from [go.dev/dl/](https://go.dev/dl/)
* follow instructions from [go.dev/doc/install](https://go.dev/doc/install)
* recommended for MacOS [formulae.brew.sh/formula/go](https://formulae.brew.sh/formula/go)

```shell
brew install go
```

----

### Go Playground
![go-playground](img/INTRODUCTION/playground.png)<!-- .element height="500px" -->
[go.dev/play](https://go.dev/play/)


----

### Useful Go Tools
* `go build` the go compiler
* `go fmt` format your go code
* `go get` load/installs dependencies
* `go run` compiles and executes go code directly
* `go test` compiles and executes go testcases
* `go mod [download|edit|graph|init|tidy|vendor|verify|why]` manages modules
* `go generate` generates code from source files
* `go version` determines the go version used

----

### Useful Go Tools
* `go vet` finds potential errors in the application
* `go bug` report error
* `go doc` standard tool to view documentation and source code
* `go env` set and read go environment variables
* `go fix` makes adjustments to new go versions
* `go tool` lists tools

----

### Set of Keywords

```
break, case, chan, const, continue
default, defer, else, fallthrough
for, func, go, goto, if, import
interface, map, package, range
return, select, struct, switch
type, var
```

----

### Base Types
* bool
* numeric types
  * int8, int16, int32, int64, int
  * uint8, uint16, uint32, uint64, uint
  * float32, float64
  * complex64, complex128
  * byte (alias for uint8)
  * rune (alias for int32)
* string

----

### Functions
```go
func add(x int, y int) int {
	return x + y
}
```

```go
func add(x, y int) int {
    return x + y
}
```

----

* multiple results

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

* named return values

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```
----
* first class functions

```go
func process(sample int, fn func(int) int) int {
	return fn(sample)
}

func adder(term int) func(a int) int {
	return func(a int) int { return a + term }
}

func increment(term int) int {
	return adder(1)(term)
}

func main() {
	sample := 23
	fmt.Println(process(sample, increment),
		process(sample, adder(100)))

}
```

----

### Variables

```go
var hot, wet, far bool
var tool string
var x int
```

* with initializers

```go
var x, y int = 1, 2
var hot, wet, tool = true, false, "screwdriver"
```

* short declaration

```go
x := 3
hot, wet, tool := true, false, "screwdriver"
```

----

### Loops

```go
for i := 0; i < 10; i++ {
	sum += i
}
```

* optional init and post statements, while loop

```go
sum := 1
for sum < 1000; {
    sum += sum
}
```

* endless loop

```go
for {
}
```

----

### If statement

```go
if x < 0 {
	return "something"
}
```
```go
if a < b {
	return a
} else {
    return b
}
```

----

### Switch Statement


```go
os := runtime.GOOS
switch os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd, plan9, windows...
    fmt.Printf("%s.\n", os)
}
```

----

### Defer

* defers the execution of a function until the surrounding function returns

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

```go
...
func main() {
    file, _ := os.Open("README.md")
    defer file.Close()
    buffer := make([]byte, 1024)
    bytesRead, _ := file.Read(buffer)
    fmt.Printf("Content: %s\n", buffer[:bytesRead])
}
...
```

----

### Structs

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

----

### Arrays

```go
var tools [2]string
a[0] = "screwdriver"
a[1] = "hammer"

primes := [6]int{2, 3, 5, 7, 11, 13}
```

----

### Slices

```go
primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
```

----

### Tour of go

* for more topics checkout the tour of go

![go-playground](img/INTRODUCTION/tour_of_go.png)<!-- .element height="400px" -->
[go.dev/tour](https://go.dev/tour/)

----

### What we have learned
* Advantages of go
* How to install go
* overview of go tools
* overview of go keywords
* overview of base types
* `tour of go` is a good starting point

----

### Further readings
* Source Code for all slides
  * [github.com/go-gurus/go_tour_src](https://github.com/go-gurus/go_tour_src/tree/main/hello-world)
* Documentation
  * [golang.org/pkg](https://golang.org/pkg/)
* Go Packages
  * [pkg.go.dev](https://pkg.go.dev/)
* Tour of Go
  * [go.dev/tour](https://go.dev/tour/)

---
