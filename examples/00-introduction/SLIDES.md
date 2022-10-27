<!-- .slide: data-background="img/INTRODUCTION/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

# Golang for Developers

----
### Autors

|                                                                |                                                                                                                                                                                                                                                                                    |
|----------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ![Manfred](img/cc_manfred.jpg)<!-- .element height="200px" --> | <ul><li>[Manfred Dreese](https://codecentric.de)<br />[dreese.de](https://dreese.de)<br />[github/i78](https://github.com/i78)<br />[twitter/dem9d](https://twitter.com/dem9d)<br /></li></ul>                                                                                     |
| ![grohmio](img/cc_grohmio.png)<!-- .element height="200px" --> | <ul><li>[Andreas Grohmann](https://codecentric.de)<br />[grohm.io](https://grohm.io)<br />[github/grohmio](https://github.com/grohmio)<br />[twitter/grohmeo](https://twitter.com/grohmeo)<br />[stackoverflow/grohmio](https://stackoverflow.com/users/6654539/grohmio)</li></ul> |


----

## Introduction
Why should you look at the GO programming language?
Go takes a lot of getting used to and is reduced to the essentials.
You can build anything from small helper scripts for pipelines to microservices to highly complex monoliths with GO.
The advantages of Go are a strong standardization of the tools,
a high applicability of the binaries and a high performance of the build process.
We will discuss further advantages of Go in more detail.

----

## Why Go?
* small binaries
* fast build times
* multi platform builds
* binaries could be used in scratch image
* small set of keywords (25)
* parallelism
* completely downward compatible

----

## Install Golang Linux, Windows, MacOS
* download your package from [go.dev/dl/](https://go.dev/dl/)
* follow instructions from [go.dev/doc/install](https://go.dev/doc/install)
* recommendet for MacOS [formulae.brew.sh/formula/go](https://formulae.brew.sh/formula/go)
```shell
brew install go
```

----
## Useful Go Tools
* `go fmt` format your go code
* `go get` load/installs dependencies
* `go run` compiles and executes go code directly
* `go test` compiles and executes go testcases
* `go mod [download|edit|graph|init|tidy|vendor|verify|why]` manages modules
* `go generate` generates code from source files
* `go version` determines the go version used
----
## Useful Go Tools
* `go vet` finds potential errors in the application
* `go bug` report error
* `go doc` standard tool to view documentation and source code
* `go env` set and read go environment variables
* `go fix` makes adjustments to new go versions
* `go tool` lists tools

----
## Set of Keywords

```
break, case, chan, const, continue
default, defer, else, fallthrough
for, func, go, goto, if, import
interface, map, package, range
return, select, struct, switch
type, var
```

----
## Base Types
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
### What we have learned
* Advantages of go
* How to install go
* overview of go tools
* overview of go keywords
* overview of base types

----
### Further readings
* Documentation
    * [golang.org/pkg/ ](https://golang.org/pkg/)
* Go Packages
    * [pkg.go.dev/](https://pkg.go.dev/)

---
